include Makefile.env
export DOCKER_TAGNAME ?= master
export KUBE_NAMESPACE ?= fybrik-system

.PHONY: all
all: generate manifests verify

.PHONY: license
license: $(TOOLBIN)/license_finder
	$(call license_go,.)

.PHONY: generate
generate: $(TOOLBIN)/controller-gen $(TOOLBIN)/json-schema-generator
	$(TOOLBIN)/json-schema-generator -r ./pkg/model/... -o charts/fybrik/files/taxonomy/
	$(TOOLBIN)/controller-gen object:headerFile=./hack/boilerplate.go.txt,year=$(shell date +%Y) paths="./..."
	$(MAKE) -C site generate

.PHONY: manifests
manifests: $(TOOLBIN)/controller-gen $(TOOLBIN)/yq
	$(TOOLBIN)/controller-gen --version
	$(TOOLBIN)/controller-gen crd output:crd:artifacts:config=charts/fybrik-crd/templates/ paths=./manager/apis/...
	$(TOOLBIN)/controller-gen webhook paths=./manager/apis/... output:stdout | \
		$(TOOLBIN)/yq eval '.metadata.annotations."cert-manager.io/inject-ca-from" |= "{{ .Release.Namespace }}/serving-cert"' - | \
		$(TOOLBIN)/yq eval '.metadata.annotations."certmanager.k8s.io/inject-ca-from" |= "{{ .Release.Namespace }}/serving-cert"' - | \
		$(TOOLBIN)/yq eval '(.metadata.name | select(. == "mutating-webhook-configuration")) = "{{ .Release.Namespace }}-mutating-webhook"' - | \
		$(TOOLBIN)/yq eval '(.metadata.name | select(. == "validating-webhook-configuration")) = "{{ .Release.Namespace }}-validating-webhook"' - | \
		$(TOOLBIN)/yq eval '(.webhooks.[].clientConfig.service.namespace) = "{{ .Release.Namespace }}"' - > charts/fybrik/files/webhook-configs.yaml

.PHONY: docker-mirror-read
docker-mirror-read:
	$(TOOLS_DIR)/docker_mirror.sh $(TOOLS_DIR)/docker_mirror.conf

.PHONY: deploy
deploy: export VALUES_FILE?=charts/fybrik/values.yaml
deploy: $(TOOLBIN)/kubectl $(TOOLBIN)/helm
	$(TOOLBIN)/kubectl create namespace $(KUBE_NAMESPACE) || true
	$(TOOLBIN)/helm install fybrik-crd charts/fybrik-crd  \
               --namespace $(KUBE_NAMESPACE) --wait --timeout 120s
	$(TOOLBIN)/helm install fybrik charts/fybrik --values $(VALUES_FILE) \
               --namespace $(KUBE_NAMESPACE) --wait --timeout 120s

pre-test: generate manifests $(TOOLBIN)/etcd $(TOOLBIN)/kube-apiserver
	mkdir -p /tmp/taxonomy
	mkdir -p /tmp/adminconfig
	cp charts/fybrik/files/taxonomy/*.json /tmp/taxonomy/
	cp charts/fybrik/files/adminconfig/*.rego /tmp/adminconfig/
	mkdir -p manager/testdata/unittests/basetaxonomy
	mkdir -p manager/testdata/unittests/sampletaxonomy
	cp charts/fybrik/files/taxonomy/*.json manager/testdata/unittests/basetaxonomy
	cp charts/fybrik/files/taxonomy/*.json manager/testdata/unittests/sampletaxonomy
	go run main.go taxonomy compile -o manager/testdata/unittests/sampletaxonomy/taxonomy.json \
  	-b charts/fybrik/files/taxonomy/taxonomy.json \
		$(shell find samples/taxonomy/example -type f -name '*.yaml')
	cp manager/testdata/unittests/sampletaxonomy/taxonomy.json /tmp/taxonomy/taxonomy.json

.PHONY: test
test: export MODULES_NAMESPACE?=fybrik-blueprints
test: export CONTROLLER_NAMESPACE?=fybrik-system
test: pre-test
	go test -v ./...

.PHONY: run-integration-tests
run-integration-tests: export DOCKER_HOSTNAME?=localhost:5000
run-integration-tests: export DOCKER_NAMESPACE?=fybrik-system
run-integration-tests: export VALUES_FILE=charts/fybrik/integration-tests.values.yaml
run-integration-tests:
	$(MAKE) kind
	$(MAKE) cluster-prepare
	$(MAKE) docker-build docker-push
	$(MAKE) -C test/services docker-build docker-push
	$(MAKE) cluster-prepare-wait
	$(MAKE) deploy
	$(MAKE) configure-vault
	$(MAKE) -C modules helm
	$(MAKE) -C modules helm-uninstall # Uninstalls the deployed tests from previous command
	$(MAKE) -C pkg/helm test
	$(MAKE) -C manager run-integration-tests
	$(MAKE) -C modules test

.PHONY: run-notebook-tests
run-notebook-tests: export DOCKER_HOSTNAME?=localhost:5000
run-notebook-tests: export DOCKER_NAMESPACE?=fybrik-system
run-notebook-tests: export VALUES_FILE=charts/fybrik/notebook-tests.values.yaml
run-notebook-tests:
	$(MAKE) kind
	$(MAKE) cluster-prepare
	$(MAKE) docker-build docker-push
	$(MAKE) -C test/services docker-build docker-push
	$(MAKE) cluster-prepare-wait
	$(MAKE) deploy
	$(MAKE) configure-vault
	$(MAKE) -C manager run-notebook-tests

.PHONY: cluster-prepare
cluster-prepare:
	$(MAKE) -C third_party/cert-manager deploy
	$(MAKE) -C third_party/vault deploy
	$(MAKE) -C third_party/datashim deploy

.PHONY: cluster-prepare-wait
cluster-prepare-wait:
	$(MAKE) -C third_party/datashim deploy-wait
	$(MAKE) -C third_party/vault deploy-wait

# Build only the docker images needed for integration testing
.PHONY: docker-minimal-it
docker-minimal-it:
	$(MAKE) -C manager docker-build docker-push
	$(MAKE) -C test/services docker-build docker-push

.PHONY: docker-build
docker-build:
	$(MAKE) -C manager docker-build
	$(MAKE) -C connectors docker-build

.PHONY: docker-push
docker-push:
	$(MAKE) -C manager docker-push
	$(MAKE) -C connectors docker-push

DOCKER_PUBLIC_HOSTNAME ?= ghcr.io
DOCKER_PUBLIC_NAMESPACE ?= fybrik
DOCKER_PUBLIC_TAGNAME ?= master

DOCKER_PUBLIC_NAMES := \
	manager \
	katalog-connector \
	opa-connector

define do-docker-retag-and-push-public
	for name in ${DOCKER_PUBLIC_NAMES}; do \
		docker tag ${DOCKER_HOSTNAME}/${DOCKER_NAMESPACE}/$$name:${DOCKER_TAGNAME} ${DOCKER_PUBLIC_HOSTNAME}/${DOCKER_PUBLIC_NAMESPACE}/$$name:${DOCKER_PUBLIC_TAGNAME}; \
	done
	DOCKER_HOSTNAME=${DOCKER_PUBLIC_HOSTNAME} DOCKER_NAMESPACE=${DOCKER_PUBLIC_NAMESPACE} DOCKER_TAGNAME=${DOCKER_PUBLIC_TAGNAME} $(MAKE) docker-push
endef

.PHONY: docker-retag-and-push-public
docker-retag-and-push-public:
	$(call do-docker-retag-and-push-public)

.PHONY: helm-push-public
helm-push-public:
	DOCKER_HOSTNAME=${DOCKER_PUBLIC_HOSTNAME} DOCKER_NAMESPACE=${DOCKER_PUBLIC_NAMESPACE} DOCKER_TAGNAME=${DOCKER_PUBLIC_TAGNAME} make -C modules helm-chart-push

.PHONY: save-images
save-images:
	docker save -o images.tar ${DOCKER_HOSTNAME}/${DOCKER_NAMESPACE}/manager:${DOCKER_TAGNAME} \
		${DOCKER_HOSTNAME}/${DOCKER_NAMESPACE}/katalog-connector:${DOCKER_TAGNAME} \
		${DOCKER_HOSTNAME}/${DOCKER_NAMESPACE}/opa-connector:${DOCKER_TAGNAME}

include hack/make-rules/tools.mk
include hack/make-rules/verify.mk
include hack/make-rules/cluster.mk
include hack/make-rules/vault.mk

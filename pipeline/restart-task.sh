#!/bin/bash

if [[ -z $1 ]]; then
    echo "Please specify a regex for the taskrun to restart

e.g.

./restart-task.sh run-integration-tests"
    exit 0
fi

echo $1

set -x
for i in $(kubectl get taskrun --no-headers | grep $1 | grep -v "restart-" | cut -d' ' -f1); do
  kubectl get taskrun $i -o yaml > /tmp/taskrun-$i.yaml
  sed -i.bak "s|  name: $i|  generateName: restart-$1|g" /tmp/taskrun-$i.yaml
  sed -i.bak "s|  namespace: .*||g" /tmp/taskrun-$i.yaml
  set +e
  sed -i.bak "s|pipeline.tekton.dev/affinity-assistant: .*||g" /tmp/taskrun-$i.yaml
  set -e
  kubectl create -f /tmp/taskrun-$i.yaml
done


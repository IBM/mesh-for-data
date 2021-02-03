// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.1
// source: data_store.proto

package connectors

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DataStore_DataStoreType int32

const (
	DataStore_UNKNOWN DataStore_DataStoreType = 0
	DataStore_LOCAL   DataStore_DataStoreType = 1
	DataStore_S3      DataStore_DataStoreType = 2
	DataStore_DB2     DataStore_DataStoreType = 3
	DataStore_KAFKA   DataStore_DataStoreType = 4
)

// Enum value maps for DataStore_DataStoreType.
var (
	DataStore_DataStoreType_name = map[int32]string{
		0: "UNKNOWN",
		1: "LOCAL",
		2: "S3",
		3: "DB2",
		4: "KAFKA",
	}
	DataStore_DataStoreType_value = map[string]int32{
		"UNKNOWN": 0,
		"LOCAL":   1,
		"S3":      2,
		"DB2":     3,
		"KAFKA":   4,
	}
)

func (x DataStore_DataStoreType) Enum() *DataStore_DataStoreType {
	p := new(DataStore_DataStoreType)
	*p = x
	return p
}

func (x DataStore_DataStoreType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataStore_DataStoreType) Descriptor() protoreflect.EnumDescriptor {
	return file_data_store_proto_enumTypes[0].Descriptor()
}

func (DataStore_DataStoreType) Type() protoreflect.EnumType {
	return &file_data_store_proto_enumTypes[0]
}

func (x DataStore_DataStoreType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataStore_DataStoreType.Descriptor instead.
func (DataStore_DataStoreType) EnumDescriptor() ([]byte, []int) {
	return file_data_store_proto_rawDescGZIP(), []int{3, 0}
}

type Db2DataStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url      string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Database string `protobuf:"bytes,2,opt,name=database,proto3" json:"database,omitempty"`
	Table    string `protobuf:"bytes,3,opt,name=table,proto3" json:"table,omitempty"` // reformat to SCHEMA.TABLE struct
	Port     string `protobuf:"bytes,4,opt,name=port,proto3" json:"port,omitempty"`
	Ssl      string `protobuf:"bytes,5,opt,name=ssl,proto3" json:"ssl,omitempty"` //Note that bool value if set to "false" does not appear in the struct at all
}

func (x *Db2DataStore) Reset() {
	*x = Db2DataStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Db2DataStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Db2DataStore) ProtoMessage() {}

func (x *Db2DataStore) ProtoReflect() protoreflect.Message {
	mi := &file_data_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Db2DataStore.ProtoReflect.Descriptor instead.
func (*Db2DataStore) Descriptor() ([]byte, []int) {
	return file_data_store_proto_rawDescGZIP(), []int{0}
}

func (x *Db2DataStore) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Db2DataStore) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *Db2DataStore) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *Db2DataStore) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *Db2DataStore) GetSsl() string {
	if x != nil {
		return x.Ssl
	}
	return ""
}

type S3DataStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint  string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Bucket    string `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	ObjectKey string `protobuf:"bytes,3,opt,name=object_key,json=objectKey,proto3" json:"object_key,omitempty"` //can be object name or the prefix for dataset
	Region    string `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`                        // WKC does not return it, it will stay empty in our case!!!
}

func (x *S3DataStore) Reset() {
	*x = S3DataStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3DataStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3DataStore) ProtoMessage() {}

func (x *S3DataStore) ProtoReflect() protoreflect.Message {
	mi := &file_data_store_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3DataStore.ProtoReflect.Descriptor instead.
func (*S3DataStore) Descriptor() ([]byte, []int) {
	return file_data_store_proto_rawDescGZIP(), []int{1}
}

func (x *S3DataStore) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *S3DataStore) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *S3DataStore) GetObjectKey() string {
	if x != nil {
		return x.ObjectKey
	}
	return ""
}

func (x *S3DataStore) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

type KafkaDataStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopicName             string `protobuf:"bytes,1,opt,name=topic_name,json=topicName,proto3" json:"topic_name,omitempty"`
	BootstrapServers      string `protobuf:"bytes,2,opt,name=bootstrap_servers,json=bootstrapServers,proto3" json:"bootstrap_servers,omitempty"`
	SchemaRegistry        string `protobuf:"bytes,3,opt,name=schema_registry,json=schemaRegistry,proto3" json:"schema_registry,omitempty"`
	KeyDeserializer       string `protobuf:"bytes,4,opt,name=key_deserializer,json=keyDeserializer,proto3" json:"key_deserializer,omitempty"`
	ValueDeserializer     string `protobuf:"bytes,5,opt,name=value_deserializer,json=valueDeserializer,proto3" json:"value_deserializer,omitempty"`
	SecurityProtocol      string `protobuf:"bytes,6,opt,name=security_protocol,json=securityProtocol,proto3" json:"security_protocol,omitempty"`
	SaslMechanism         string `protobuf:"bytes,7,opt,name=sasl_mechanism,json=saslMechanism,proto3" json:"sasl_mechanism,omitempty"`
	SslTruststore         string `protobuf:"bytes,8,opt,name=ssl_truststore,json=sslTruststore,proto3" json:"ssl_truststore,omitempty"`
	SslTruststorePassword string `protobuf:"bytes,9,opt,name=ssl_truststore_password,json=sslTruststorePassword,proto3" json:"ssl_truststore_password,omitempty"`
}

func (x *KafkaDataStore) Reset() {
	*x = KafkaDataStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_store_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KafkaDataStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KafkaDataStore) ProtoMessage() {}

func (x *KafkaDataStore) ProtoReflect() protoreflect.Message {
	mi := &file_data_store_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KafkaDataStore.ProtoReflect.Descriptor instead.
func (*KafkaDataStore) Descriptor() ([]byte, []int) {
	return file_data_store_proto_rawDescGZIP(), []int{2}
}

func (x *KafkaDataStore) GetTopicName() string {
	if x != nil {
		return x.TopicName
	}
	return ""
}

func (x *KafkaDataStore) GetBootstrapServers() string {
	if x != nil {
		return x.BootstrapServers
	}
	return ""
}

func (x *KafkaDataStore) GetSchemaRegistry() string {
	if x != nil {
		return x.SchemaRegistry
	}
	return ""
}

func (x *KafkaDataStore) GetKeyDeserializer() string {
	if x != nil {
		return x.KeyDeserializer
	}
	return ""
}

func (x *KafkaDataStore) GetValueDeserializer() string {
	if x != nil {
		return x.ValueDeserializer
	}
	return ""
}

func (x *KafkaDataStore) GetSecurityProtocol() string {
	if x != nil {
		return x.SecurityProtocol
	}
	return ""
}

func (x *KafkaDataStore) GetSaslMechanism() string {
	if x != nil {
		return x.SaslMechanism
	}
	return ""
}

func (x *KafkaDataStore) GetSslTruststore() string {
	if x != nil {
		return x.SslTruststore
	}
	return ""
}

func (x *KafkaDataStore) GetSslTruststorePassword() string {
	if x != nil {
		return x.SslTruststorePassword
	}
	return ""
}

type DataStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type DataStore_DataStoreType `protobuf:"varint,1,opt,name=type,proto3,enum=connectors.DataStore_DataStoreType" json:"type,omitempty"`
	Name string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` //for auditing and readability. Can be same as location type or can have more info if availble from catalog
	// oneof location {   // should have been oneof but for technical rasons, a problem to translate it to JSON, we remove the oneof for now
	//should have been local, db2, s3 without "location"  but had a problem to compile it in proto - collision with proto name DataLocationDb2
	Db2   *Db2DataStore   `protobuf:"bytes,3,opt,name=db2,proto3" json:"db2,omitempty"`
	S3    *S3DataStore    `protobuf:"bytes,4,opt,name=s3,proto3" json:"s3,omitempty"`
	Kafka *KafkaDataStore `protobuf:"bytes,5,opt,name=kafka,proto3" json:"kafka,omitempty"`
}

func (x *DataStore) Reset() {
	*x = DataStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_store_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataStore) ProtoMessage() {}

func (x *DataStore) ProtoReflect() protoreflect.Message {
	mi := &file_data_store_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataStore.ProtoReflect.Descriptor instead.
func (*DataStore) Descriptor() ([]byte, []int) {
	return file_data_store_proto_rawDescGZIP(), []int{3}
}

func (x *DataStore) GetType() DataStore_DataStoreType {
	if x != nil {
		return x.Type
	}
	return DataStore_UNKNOWN
}

func (x *DataStore) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DataStore) GetDb2() *Db2DataStore {
	if x != nil {
		return x.Db2
	}
	return nil
}

func (x *DataStore) GetS3() *S3DataStore {
	if x != nil {
		return x.S3
	}
	return nil
}

func (x *DataStore) GetKafka() *KafkaDataStore {
	if x != nil {
		return x.Kafka
	}
	return nil
}

var File_data_store_proto protoreflect.FileDescriptor

var file_data_store_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x78,
	0x0a, 0x0c, 0x44, 0x62, 0x32, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x73, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x73, 0x6c, 0x22, 0x78, 0x0a, 0x0b, 0x53, 0x33, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x22, 0x92, 0x03, 0x0a, 0x0e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61,
	0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x73, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x29, 0x0a, 0x10, 0x6b, 0x65,
	0x79, 0x5f, 0x64, 0x65, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6b, 0x65, 0x79, 0x44, 0x65, 0x73, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x12, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x64,
	0x65, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x44, 0x65, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x61, 0x73, 0x6c, 0x5f, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e,
	0x69, 0x73, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x61, 0x73, 0x6c, 0x4d,
	0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x73, 0x6c, 0x5f,
	0x74, 0x72, 0x75, 0x73, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x73, 0x73, 0x6c, 0x54, 0x72, 0x75, 0x73, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12,
	0x36, 0x0a, 0x17, 0x73, 0x73, 0x6c, 0x5f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x15, 0x73, 0x73, 0x6c, 0x54, 0x72, 0x75, 0x73, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xa4, 0x02, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x03, 0x64, 0x62, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x44, 0x62, 0x32,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x03, 0x64, 0x62, 0x32, 0x12, 0x27,
	0x0a, 0x02, 0x73, 0x33, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x53, 0x33, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x52, 0x02, 0x73, 0x33, 0x12, 0x30, 0x0a, 0x05, 0x6b, 0x61, 0x66, 0x6b, 0x61,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x73, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x52, 0x05, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x22, 0x43, 0x0a, 0x0d, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x4f, 0x43, 0x41, 0x4c,
	0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x53, 0x33, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x42,
	0x32, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x4b, 0x41, 0x46, 0x4b, 0x41, 0x10, 0x04, 0x42, 0x0d,
	0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x74, 0x6d, 0x65, 0x73, 0x68, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_store_proto_rawDescOnce sync.Once
	file_data_store_proto_rawDescData = file_data_store_proto_rawDesc
)

func file_data_store_proto_rawDescGZIP() []byte {
	file_data_store_proto_rawDescOnce.Do(func() {
		file_data_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_store_proto_rawDescData)
	})
	return file_data_store_proto_rawDescData
}

var file_data_store_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_data_store_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_data_store_proto_goTypes = []interface{}{
	(DataStore_DataStoreType)(0), // 0: connectors.DataStore.DataStoreType
	(*Db2DataStore)(nil),         // 1: connectors.Db2DataStore
	(*S3DataStore)(nil),          // 2: connectors.S3DataStore
	(*KafkaDataStore)(nil),       // 3: connectors.KafkaDataStore
	(*DataStore)(nil),            // 4: connectors.DataStore
}
var file_data_store_proto_depIdxs = []int32{
	0, // 0: connectors.DataStore.type:type_name -> connectors.DataStore.DataStoreType
	1, // 1: connectors.DataStore.db2:type_name -> connectors.Db2DataStore
	2, // 2: connectors.DataStore.s3:type_name -> connectors.S3DataStore
	3, // 3: connectors.DataStore.kafka:type_name -> connectors.KafkaDataStore
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_data_store_proto_init() }
func file_data_store_proto_init() {
	if File_data_store_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Db2DataStore); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_data_store_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3DataStore); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_data_store_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KafkaDataStore); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_data_store_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataStore); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_data_store_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_store_proto_goTypes,
		DependencyIndexes: file_data_store_proto_depIdxs,
		EnumInfos:         file_data_store_proto_enumTypes,
		MessageInfos:      file_data_store_proto_msgTypes,
	}.Build()
	File_data_store_proto = out.File
	file_data_store_proto_rawDesc = nil
	file_data_store_proto_goTypes = nil
	file_data_store_proto_depIdxs = nil
}

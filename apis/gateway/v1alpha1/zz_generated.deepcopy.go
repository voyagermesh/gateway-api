//go:build !ignore_autogenerated

/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/gateway-api/apis/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaRoute) DeepCopyInto(out *KafkaRoute) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaRoute.
func (in *KafkaRoute) DeepCopy() *KafkaRoute {
	if in == nil {
		return nil
	}
	out := new(KafkaRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KafkaRoute) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaRouteList) DeepCopyInto(out *KafkaRouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KafkaRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaRouteList.
func (in *KafkaRouteList) DeepCopy() *KafkaRouteList {
	if in == nil {
		return nil
	}
	out := new(KafkaRouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KafkaRouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaRouteRule) DeepCopyInto(out *KafkaRouteRule) {
	*out = *in
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = make([]RouteFilter, len(*in))
		copy(*out, *in)
	}
	if in.BackendRefs != nil {
		in, out := &in.BackendRefs, &out.BackendRefs
		*out = make([]v1.BackendRef, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaRouteRule.
func (in *KafkaRouteRule) DeepCopy() *KafkaRouteRule {
	if in == nil {
		return nil
	}
	out := new(KafkaRouteRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaRouteSpec) DeepCopyInto(out *KafkaRouteSpec) {
	*out = *in
	in.CommonRouteSpec.DeepCopyInto(&out.CommonRouteSpec)
	if in.Hostnames != nil {
		in, out := &in.Hostnames, &out.Hostnames
		*out = make([]v1.Hostname, len(*in))
		copy(*out, *in)
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]KafkaRouteRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaRouteSpec.
func (in *KafkaRouteSpec) DeepCopy() *KafkaRouteSpec {
	if in == nil {
		return nil
	}
	out := new(KafkaRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaRouteStatus) DeepCopyInto(out *KafkaRouteStatus) {
	*out = *in
	in.RouteStatus.DeepCopyInto(&out.RouteStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaRouteStatus.
func (in *KafkaRouteStatus) DeepCopy() *KafkaRouteStatus {
	if in == nil {
		return nil
	}
	out := new(KafkaRouteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MSSQLServerRoute) DeepCopyInto(out *MSSQLServerRoute) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MSSQLServerRoute.
func (in *MSSQLServerRoute) DeepCopy() *MSSQLServerRoute {
	if in == nil {
		return nil
	}
	out := new(MSSQLServerRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MSSQLServerRoute) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MSSQLServerRouteList) DeepCopyInto(out *MSSQLServerRouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MSSQLServerRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MSSQLServerRouteList.
func (in *MSSQLServerRouteList) DeepCopy() *MSSQLServerRouteList {
	if in == nil {
		return nil
	}
	out := new(MSSQLServerRouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MSSQLServerRouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MSSQLServerRouteRule) DeepCopyInto(out *MSSQLServerRouteRule) {
	*out = *in
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = make([]RouteFilter, len(*in))
		copy(*out, *in)
	}
	if in.BackendRefs != nil {
		in, out := &in.BackendRefs, &out.BackendRefs
		*out = make([]v1.BackendRef, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MSSQLServerRouteRule.
func (in *MSSQLServerRouteRule) DeepCopy() *MSSQLServerRouteRule {
	if in == nil {
		return nil
	}
	out := new(MSSQLServerRouteRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MSSQLServerRouteSpec) DeepCopyInto(out *MSSQLServerRouteSpec) {
	*out = *in
	in.CommonRouteSpec.DeepCopyInto(&out.CommonRouteSpec)
	if in.Hostnames != nil {
		in, out := &in.Hostnames, &out.Hostnames
		*out = make([]v1.Hostname, len(*in))
		copy(*out, *in)
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]MSSQLServerRouteRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Telemetry = in.Telemetry
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MSSQLServerRouteSpec.
func (in *MSSQLServerRouteSpec) DeepCopy() *MSSQLServerRouteSpec {
	if in == nil {
		return nil
	}
	out := new(MSSQLServerRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MSSQLServerRouteStatus) DeepCopyInto(out *MSSQLServerRouteStatus) {
	*out = *in
	in.RouteStatus.DeepCopyInto(&out.RouteStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MSSQLServerRouteStatus.
func (in *MSSQLServerRouteStatus) DeepCopy() *MSSQLServerRouteStatus {
	if in == nil {
		return nil
	}
	out := new(MSSQLServerRouteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBRoute) DeepCopyInto(out *MongoDBRoute) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBRoute.
func (in *MongoDBRoute) DeepCopy() *MongoDBRoute {
	if in == nil {
		return nil
	}
	out := new(MongoDBRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongoDBRoute) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBRouteList) DeepCopyInto(out *MongoDBRouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MongoDBRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBRouteList.
func (in *MongoDBRouteList) DeepCopy() *MongoDBRouteList {
	if in == nil {
		return nil
	}
	out := new(MongoDBRouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongoDBRouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBRouteRule) DeepCopyInto(out *MongoDBRouteRule) {
	*out = *in
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = make([]RouteFilter, len(*in))
		copy(*out, *in)
	}
	if in.BackendRefs != nil {
		in, out := &in.BackendRefs, &out.BackendRefs
		*out = make([]v1.BackendRef, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBRouteRule.
func (in *MongoDBRouteRule) DeepCopy() *MongoDBRouteRule {
	if in == nil {
		return nil
	}
	out := new(MongoDBRouteRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBRouteSpec) DeepCopyInto(out *MongoDBRouteSpec) {
	*out = *in
	in.CommonRouteSpec.DeepCopyInto(&out.CommonRouteSpec)
	if in.Hostnames != nil {
		in, out := &in.Hostnames, &out.Hostnames
		*out = make([]v1.Hostname, len(*in))
		copy(*out, *in)
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]MongoDBRouteRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBRouteSpec.
func (in *MongoDBRouteSpec) DeepCopy() *MongoDBRouteSpec {
	if in == nil {
		return nil
	}
	out := new(MongoDBRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBRouteStatus) DeepCopyInto(out *MongoDBRouteStatus) {
	*out = *in
	in.RouteStatus.DeepCopyInto(&out.RouteStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBRouteStatus.
func (in *MongoDBRouteStatus) DeepCopy() *MongoDBRouteStatus {
	if in == nil {
		return nil
	}
	out := new(MongoDBRouteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLRoute) DeepCopyInto(out *MySQLRoute) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLRoute.
func (in *MySQLRoute) DeepCopy() *MySQLRoute {
	if in == nil {
		return nil
	}
	out := new(MySQLRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MySQLRoute) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLRouteList) DeepCopyInto(out *MySQLRouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MySQLRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLRouteList.
func (in *MySQLRouteList) DeepCopy() *MySQLRouteList {
	if in == nil {
		return nil
	}
	out := new(MySQLRouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MySQLRouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLRouteRule) DeepCopyInto(out *MySQLRouteRule) {
	*out = *in
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = make([]RouteFilter, len(*in))
		copy(*out, *in)
	}
	if in.BackendRefs != nil {
		in, out := &in.BackendRefs, &out.BackendRefs
		*out = make([]v1.BackendRef, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLRouteRule.
func (in *MySQLRouteRule) DeepCopy() *MySQLRouteRule {
	if in == nil {
		return nil
	}
	out := new(MySQLRouteRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLRouteSpec) DeepCopyInto(out *MySQLRouteSpec) {
	*out = *in
	in.CommonRouteSpec.DeepCopyInto(&out.CommonRouteSpec)
	if in.Hostnames != nil {
		in, out := &in.Hostnames, &out.Hostnames
		*out = make([]v1.Hostname, len(*in))
		copy(*out, *in)
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]MySQLRouteRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Telemetry = in.Telemetry
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLRouteSpec.
func (in *MySQLRouteSpec) DeepCopy() *MySQLRouteSpec {
	if in == nil {
		return nil
	}
	out := new(MySQLRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLRouteStatus) DeepCopyInto(out *MySQLRouteStatus) {
	*out = *in
	in.RouteStatus.DeepCopyInto(&out.RouteStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLRouteStatus.
func (in *MySQLRouteStatus) DeepCopy() *MySQLRouteStatus {
	if in == nil {
		return nil
	}
	out := new(MySQLRouteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OtelProvider) DeepCopyInto(out *OtelProvider) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OtelProvider.
func (in *OtelProvider) DeepCopy() *OtelProvider {
	if in == nil {
		return nil
	}
	out := new(OtelProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresRoute) DeepCopyInto(out *PostgresRoute) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresRoute.
func (in *PostgresRoute) DeepCopy() *PostgresRoute {
	if in == nil {
		return nil
	}
	out := new(PostgresRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresRoute) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresRouteList) DeepCopyInto(out *PostgresRouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PostgresRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresRouteList.
func (in *PostgresRouteList) DeepCopy() *PostgresRouteList {
	if in == nil {
		return nil
	}
	out := new(PostgresRouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresRouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresRouteRule) DeepCopyInto(out *PostgresRouteRule) {
	*out = *in
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = make([]RouteFilter, len(*in))
		copy(*out, *in)
	}
	if in.BackendRefs != nil {
		in, out := &in.BackendRefs, &out.BackendRefs
		*out = make([]v1.BackendRef, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresRouteRule.
func (in *PostgresRouteRule) DeepCopy() *PostgresRouteRule {
	if in == nil {
		return nil
	}
	out := new(PostgresRouteRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresRouteSpec) DeepCopyInto(out *PostgresRouteSpec) {
	*out = *in
	in.CommonRouteSpec.DeepCopyInto(&out.CommonRouteSpec)
	if in.Hostnames != nil {
		in, out := &in.Hostnames, &out.Hostnames
		*out = make([]v1.Hostname, len(*in))
		copy(*out, *in)
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]PostgresRouteRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Telemetry = in.Telemetry
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSSocketConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresRouteSpec.
func (in *PostgresRouteSpec) DeepCopy() *PostgresRouteSpec {
	if in == nil {
		return nil
	}
	out := new(PostgresRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresRouteStatus) DeepCopyInto(out *PostgresRouteStatus) {
	*out = *in
	in.RouteStatus.DeepCopyInto(&out.RouteStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresRouteStatus.
func (in *PostgresRouteStatus) DeepCopy() *PostgresRouteStatus {
	if in == nil {
		return nil
	}
	out := new(PostgresRouteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisRoute) DeepCopyInto(out *RedisRoute) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisRoute.
func (in *RedisRoute) DeepCopy() *RedisRoute {
	if in == nil {
		return nil
	}
	out := new(RedisRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisRoute) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisRouteList) DeepCopyInto(out *RedisRouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedisRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisRouteList.
func (in *RedisRouteList) DeepCopy() *RedisRouteList {
	if in == nil {
		return nil
	}
	out := new(RedisRouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisRouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisRouteRule) DeepCopyInto(out *RedisRouteRule) {
	*out = *in
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = make([]RouteFilter, len(*in))
		copy(*out, *in)
	}
	if in.BackendRefs != nil {
		in, out := &in.BackendRefs, &out.BackendRefs
		*out = make([]v1.BackendRef, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisRouteRule.
func (in *RedisRouteRule) DeepCopy() *RedisRouteRule {
	if in == nil {
		return nil
	}
	out := new(RedisRouteRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisRouteSpec) DeepCopyInto(out *RedisRouteSpec) {
	*out = *in
	in.CommonRouteSpec.DeepCopyInto(&out.CommonRouteSpec)
	if in.Hostnames != nil {
		in, out := &in.Hostnames, &out.Hostnames
		*out = make([]v1.Hostname, len(*in))
		copy(*out, *in)
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]RedisRouteRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AuthSecret != nil {
		in, out := &in.AuthSecret, &out.AuthSecret
		*out = new(corev1.SecretReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisRouteSpec.
func (in *RedisRouteSpec) DeepCopy() *RedisRouteSpec {
	if in == nil {
		return nil
	}
	out := new(RedisRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisRouteStatus) DeepCopyInto(out *RedisRouteStatus) {
	*out = *in
	in.RouteStatus.DeepCopyInto(&out.RouteStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisRouteStatus.
func (in *RedisRouteStatus) DeepCopy() *RedisRouteStatus {
	if in == nil {
		return nil
	}
	out := new(RedisRouteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteFilter) DeepCopyInto(out *RouteFilter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteFilter.
func (in *RouteFilter) DeepCopy() *RouteFilter {
	if in == nil {
		return nil
	}
	out := new(RouteFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSSocketConfig) DeepCopyInto(out *TLSSocketConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSSocketConfig.
func (in *TLSSocketConfig) DeepCopy() *TLSSocketConfig {
	if in == nil {
		return nil
	}
	out := new(TLSSocketConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TelemetryConfig) DeepCopyInto(out *TelemetryConfig) {
	*out = *in
	out.Provider = in.Provider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TelemetryConfig.
func (in *TelemetryConfig) DeepCopy() *TelemetryConfig {
	if in == nil {
		return nil
	}
	out := new(TelemetryConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TelemetryRefence) DeepCopyInto(out *TelemetryRefence) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TelemetryRefence.
func (in *TelemetryRefence) DeepCopy() *TelemetryRefence {
	if in == nil {
		return nil
	}
	out := new(TelemetryRefence)
	in.DeepCopyInto(out)
	return out
}

/*
Copyright 2023.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// EtcdPerfSpec defines the desired state of EtcdPerf
type EtcdPerfSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of EtcdPerf. Edit etcdperf_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// EtcdPerfStatus defines the observed state of EtcdPerf
type EtcdPerfStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// EtcdPerf is the Schema for the etcdperves API
type EtcdPerf struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EtcdPerfSpec   `json:"spec,omitempty"`
	Status EtcdPerfStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// EtcdPerfList contains a list of EtcdPerf
type EtcdPerfList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EtcdPerf `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EtcdPerf{}, &EtcdPerfList{})
}

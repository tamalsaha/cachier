/*
Copyright 2018 The Knative Authors

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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	"knative.dev/pkg/apis/duck"
	"knative.dev/pkg/kmeta"
)

// PodSpecable is implemented by types containing a PodTemplateSpec
// in the manner of ReplicaSet, Deployment, DaemonSet, StatefulSet.
type PodSpecable corev1.PodTemplateSpec

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type WithPod struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec WithPodSpec `json:"spec,omitempty"`
}

type WithPodSpec struct {
	Template PodSpecable `json:"template,omitempty"`
}

// Ensure WithPod satisfies apis.Listable
var _ apis.Listable = (*WithPod)(nil)

// Ensure WithPod satisfies apis.Listable
var _ kmeta.OwnerRefable = (*WithPod)(nil)

// TODO(mattmoor): Move to tests
var _ duck.Populatable = (*WithPod)(nil)
var _ duck.Implementable = (*PodSpecable)(nil)

// GetFullType implements duck.Implementable
func (_ *PodSpecable) GetFullType() duck.Populatable {
	return &WithPod{}
}

func (t *WithPod) GetGroupVersionKind() schema.GroupVersionKind {
	return t.TypeMeta.GroupVersionKind()
}

// Populate implements duck.Populatable
func (t *WithPod) Populate() {
	t.Spec.Template = PodSpecable{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"foo": "bar",
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{{
				Name:  "container-name",
				Image: "container-image:latest",
			}},
		},
	}
}

// GetListType implements apis.Listable
func (r *WithPod) GetListType() runtime.Object {
	return &WithPodList{}
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WithPodList is a list of WithPod resources
type WithPodList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []WithPod `json:"items"`
}

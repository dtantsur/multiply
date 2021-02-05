/*


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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MultiplySpec defines the desired state of Multiply
type MultiplySpec struct {
	Operand1 int32 `json:"operand1"`
	Operand2 int32 `json:"operand2"`
}

// MultiplyStatus defines the observed state of Multiply
type MultiplyStatus struct {
	Result   int32 `json:"result"`
	Overflow bool  `json:"overflow"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// Multiply is the Schema for the multiplies API
type Multiply struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MultiplySpec   `json:"spec,omitempty"`
	Status MultiplyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MultiplyList contains a list of Multiply
type MultiplyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Multiply `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Multiply{}, &MultiplyList{})
}

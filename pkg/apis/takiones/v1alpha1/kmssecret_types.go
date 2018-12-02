package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KmsSecretSpec defines the desired state of KmsSecret
type KmsSecretSpec struct {
	Data map[string][]byte `json:"data"`

	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
}

// KmsSecretStatus defines the observed state of KmsSecret
type KmsSecretStatus struct {
	Created                 bool  `json:"created"`
	LastModificationApplied int64 `json:"lastModificatonApplied"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KmsSecret is the Schema for the kmssecrets API
// +k8s:openapi-gen=true
type KmsSecret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KmsSecretSpec   `json:"spec,omitempty"`
	Status KmsSecretStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KmsSecretList contains a list of KmsSecret
type KmsSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KmsSecret `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KmsSecret{}, &KmsSecretList{})
}

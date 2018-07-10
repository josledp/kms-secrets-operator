package v1alpha1

import (
	"encoding/json"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KmsSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []KmsSecret `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KmsSecret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              KmsSecretSpec   `json:"spec"`
	Status            KmsSecretStatus `json:"status,omitempty"`
}

type KmsSecretSpec struct {
	Data map[string][]byte `json:"data"`
}
type KmsSecretStatus struct {
	Created                 bool      `json:"created"`
	LastModificationApplied time.Time `json:"lastModificatonApplied"`
}

func (k *KmsSecret) GetVersion() string {
	version, _ := json.Marshal(k.Spec)
	return string(version)
}

func (KmsSecret) GetVersionAnnotationName() string {
	return "kms-secrets-operator/last-version-applied"
}

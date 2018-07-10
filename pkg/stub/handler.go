package stub

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/josledp/kms-secrets-operator/pkg/apis/takiones/v1alpha1"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

func NewHandler() sdk.Handler {
	region := "us-east-1"
	awscfg := aws.NewConfig()
	awscfg.Region = &region
	return &Handler{
		kmsSvc: kms.New(session.New(), awscfg),
	}
}

type Handler struct {
	kmsSvc *kms.KMS
	// Fill me
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	switch o := event.Object.(type) {
	case *v1alpha1.KmsSecret:
		//First handle de easy event :)
		if event.Deleted {
			err := sdk.Delete(getSecretBase(o))
			if err != nil && !errors.IsNotFound(err) {
				return err
			}
			return nil
		}

		//Check if the kms secret has been modified
		baseObj := getSecretBase(o)
		err := sdk.Get(baseObj)
		var rv string
		if err != nil && !errors.IsNotFound(err) {
			logrus.Errorf("unable to get current secret: %v", err)
			return err
		} else if err == nil {
			rv, _ = baseObj.Annotations[o.GetVersionAnnotationName()]
		}
		if rv == o.GetVersion() {
			return nil
		}

		// We have to create or update it
		secret, err := h.getSecret(o)
		if err != nil {
			logrus.Errorf("unable to create secret: %v", err)
			return fmt.Errorf("unable to create secret: %v", err)
		}
		err = sdk.Create(secret)
		if err != nil && !errors.IsAlreadyExists(err) {
			logrus.Errorf("failed to create secret: %v", err)
			return err
		} else if errors.IsAlreadyExists(err) {
			err = sdk.Update(secret)
			if err != nil {
				logrus.Errorf("failed to update secret: %v", err)
				return err
			}
		}

		//update kmsSecret status
		o.Status = v1alpha1.KmsSecretStatus{
			Created:                 true,
			LastModificationApplied: time.Now(),
		}
		err = sdk.Update(o)
		if err != nil {
			logrus.Errorf("error updating kmsSecret status: %v", err)
			return err
		}

	}
	return nil
}

func (h *Handler) getSecret(cr *v1alpha1.KmsSecret) (*corev1.Secret, error) {
	secret := getSecretBase(cr)
	data := make(map[string][]byte)

	for k, v := range cr.Spec.Data {
		input := &kms.DecryptInput{
			CiphertextBlob: v,
		}
		output, err := h.kmsSvc.Decrypt(input)
		if err != nil {
			return nil, err
		}
		data[k] = output.Plaintext
	}
	secret.Data = data

	secret.Annotations = map[string]string{
		cr.GetVersionAnnotationName(): cr.GetVersion(),
	}
	return secret, nil
}

func getSecretBase(cr *v1alpha1.KmsSecret) *corev1.Secret {
	secret := &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.ObjectMeta.Name,
			Namespace: cr.ObjectMeta.Namespace,
			Labels:    cr.ObjectMeta.Labels,
		},
	}

	secret.SetOwnerReferences([]metav1.OwnerReference{
		*metav1.NewControllerRef(cr, schema.GroupVersionKind{
			Group:   v1alpha1.SchemeGroupVersion.Group,
			Version: v1alpha1.SchemeGroupVersion.Version,
			Kind:    "KmsSecret",
		}),
	})
	return secret
}

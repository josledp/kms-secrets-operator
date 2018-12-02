package controller

import (
	"github.com/josledp/kms-secrets-operator/pkg/controller/kmssecret"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, kmssecret.Add)
}

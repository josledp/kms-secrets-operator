all: build_image upload

image := "quay.io/josledp/kms-secrets-operator"
version := $(shell git describe --tag)

build_image:
	operator-sdk generate k8s
	operator-sdk build $(image):$(version)

upload:
	docker login quay.io
	docker push $(image)

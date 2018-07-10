all: build upload

image := "quay.io/josledp/kms-secrets-operator"
version := $(shell git describe --tag)

build:
	operator-sdk build $(image):$(version)

upload:
	docker login quay.io
	docker push $(image)

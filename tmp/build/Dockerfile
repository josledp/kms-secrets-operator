FROM alpine:3.6

RUN apk add --no-cache ca-certificates

RUN adduser -D kms-secrets-operator
USER kms-secrets-operator


ADD tmp/_output/bin/kms-secrets-operator /usr/local/bin/kms-secrets-operator

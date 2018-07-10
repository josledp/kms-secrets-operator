kms-secrets-operator is a Kubernetes operator that allows creating secrets from KMS encrypted data. It makes use of https://github.com/operator-framework/operator-sdk

kms-secrets-operator does not encrypt anything at rest (check https://kubernetes.io/docs/tasks/administer-cluster/encrypt-data/), but kms-secrets-operator allows having yaml files with the secrets encrypted, so you do not have any secret in plaintext in your repository (because you commit your Kubernetes objects yamls in a repo, right?)

After this operator has been deployed any KmsSecret object created/updated will create/update a Secret with the data decrypted with the same name/namespace.

# Setup
- Download crd.yaml, rbac.yaml and kms-secrets-operator.yaml from deploy/
- (optional) Modify the namespace where you want to deploy this operator
- Setup on kms-secrets-operator.yaml the IAM Role which can decrypt with the desired keys(if you use kube2iam/kiam, which you should!)
- kubectl apply the files

# Usage
- There is an example.yaml in the deploy/
- To encrypt data with the aws-cli:
  `aws kms encrypt --key-id "<the id of your kms key>" --plaintext "<the secret>" `

module github.com/pulumi/pulumi-godaddy

go 1.13

require (
	github.com/baronfel/pulumi-godaddy v0.0.0-20200114182609-e673654a523c
	github.com/hashicorp/terraform v0.12.17 // indirect
	github.com/n3integration/terraform-godaddy v1.7.0
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.8.0
	github.com/pulumi/pulumi-terraform v0.18.4-0.20191202134852-87cfb4dc8ae1
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)

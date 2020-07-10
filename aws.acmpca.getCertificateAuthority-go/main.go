package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/acmpca"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := acmpca.LookupCertificateAuthority(ctx, &acmpca.LookupCertificateAuthorityArgs{
			Arn: "arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}


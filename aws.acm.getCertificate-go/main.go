package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/acm"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := acm.LookupCertificate(ctx, &acm.LookupCertificateArgs{
			Domain: "tf.example.com",
			KeyTypes: []string{
				"RSA_4096",
			},
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}


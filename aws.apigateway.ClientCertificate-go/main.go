package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apigateway.NewClientCertificate(ctx, "demo", &apigateway.ClientCertificateArgs{
			Description: pulumi.String("My client certificate"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}


package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cognito"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := cognito.NewUserPool(ctx, "example", nil)
		if err != nil {
			return err
		}
		_, err = cognito.NewUserPoolDomain(ctx, "main", &cognito.UserPoolDomainArgs{
			Domain:     pulumi.String("example-domain"),
			UserPoolId: example.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}


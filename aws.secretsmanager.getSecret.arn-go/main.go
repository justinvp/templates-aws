package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/secretsmanager"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := "arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456"
		_, err := secretsmanager.LookupSecret(ctx, &secretsmanager.LookupSecretArgs{
			Arn: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}


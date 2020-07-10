package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/secretsmanager"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := secretsmanager.LookupSecretRotation(ctx, &secretsmanager.LookupSecretRotationArgs{
			SecretId: data.Aws_secretsmanager_secret.Example.Id,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}


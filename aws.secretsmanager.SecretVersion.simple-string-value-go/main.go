package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/secretsmanager"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := secretsmanager.NewSecretVersion(ctx, "example", &secretsmanager.SecretVersionArgs{
			SecretId:     pulumi.Any(aws_secretsmanager_secret.Example.Id),
			SecretString: pulumi.String("example-string-to-protect"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}


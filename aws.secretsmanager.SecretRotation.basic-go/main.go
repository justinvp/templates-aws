package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/secretsmanager"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := secretsmanager.NewSecretRotation(ctx, "example", &secretsmanager.SecretRotationArgs{
			RotationLambdaArn: pulumi.Any(aws_lambda_function.Example.Arn),
			RotationRules: &secretsmanager.SecretRotationRotationRulesArgs{
				AutomaticallyAfterDays: pulumi.Int(30),
			},
			SecretId: pulumi.Any(aws_secretsmanager_secret.Example.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}


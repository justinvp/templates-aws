package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/secretsmanager"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := secretsmanager.NewSecret(ctx, "rotation_example", &secretsmanager.SecretArgs{
			RotationLambdaArn: pulumi.Any(aws_lambda_function.Example.Arn),
			RotationRules: &secretsmanager.SecretRotationRulesArgs{
				AutomaticallyAfterDays: pulumi.Int(7),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}


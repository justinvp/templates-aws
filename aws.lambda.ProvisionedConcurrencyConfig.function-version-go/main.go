package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lambda.NewProvisionedConcurrencyConfig(ctx, "example", &lambda.ProvisionedConcurrencyConfigArgs{
			FunctionName:                    pulumi.Any(aws_lambda_function.Example.Function_name),
			ProvisionedConcurrentExecutions: pulumi.Int(1),
			Qualifier:                       pulumi.Any(aws_lambda_function.Example.Version),
		})
		if err != nil {
			return err
		}
		return nil
	})
}


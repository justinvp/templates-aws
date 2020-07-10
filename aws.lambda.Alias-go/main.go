package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lambda.NewAlias(ctx, "testLambdaAlias", &lambda.AliasArgs{
			Description:     pulumi.String("a sample description"),
			FunctionName:    pulumi.Any(aws_lambda_function.Lambda_function_test.Arn),
			FunctionVersion: pulumi.String("1"),
			RoutingConfig: &lambda.AliasRoutingConfigArgs{
				AdditionalVersionWeights: pulumi.Float64Map{
					"2": pulumi.Float64(0.5),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}


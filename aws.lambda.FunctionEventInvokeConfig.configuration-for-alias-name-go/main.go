package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lambda.NewFunctionEventInvokeConfig(ctx, "example", &lambda.FunctionEventInvokeConfigArgs{
			FunctionName: pulumi.Any(aws_lambda_alias.Example.Function_name),
			Qualifier:    pulumi.Any(aws_lambda_alias.Example.Name),
		})
		if err != nil {
			return err
		}
		return nil
	})
}


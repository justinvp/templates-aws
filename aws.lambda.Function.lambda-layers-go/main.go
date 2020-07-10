package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleLayerVersion, err := lambda.NewLayerVersion(ctx, "exampleLayerVersion", nil)
		if err != nil {
			return err
		}
		_, err = lambda.NewFunction(ctx, "exampleFunction", &lambda.FunctionArgs{
			Layers: pulumi.StringArray{
				exampleLayerVersion.Arn,
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}


package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/sagemaker"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sagemaker.NewEndpoint(ctx, "endpoint", &sagemaker.EndpointArgs{
			EndpointConfigName: pulumi.Any(aws_sagemaker_endpoint_configuration.Ec.Name),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("foo"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}


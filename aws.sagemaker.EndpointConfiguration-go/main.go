package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/sagemaker"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sagemaker.NewEndpointConfiguration(ctx, "ec", &sagemaker.EndpointConfigurationArgs{
			ProductionVariants: sagemaker.EndpointConfigurationProductionVariantArray{
				&sagemaker.EndpointConfigurationProductionVariantArgs{
					InitialInstanceCount: pulumi.Int(1),
					InstanceType:         pulumi.String("ml.t2.medium"),
					ModelName:            pulumi.Any(aws_sagemaker_model.M.Name),
					VariantName:          pulumi.String("variant-1"),
				},
			},
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


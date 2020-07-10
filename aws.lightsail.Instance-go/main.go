package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lightsail"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lightsail.NewInstance(ctx, "gitlabTest", &lightsail.InstanceArgs{
			AvailabilityZone: pulumi.String("us-east-1b"),
			BlueprintId:      pulumi.String("string"),
			BundleId:         pulumi.String("string"),
			KeyPairName:      pulumi.String("some_key_name"),
			Tags: pulumi.StringMap{
				"foo": pulumi.String("bar"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}


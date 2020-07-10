package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewSpotInstanceRequest(ctx, "cheapWorker", &ec2.SpotInstanceRequestArgs{
			Ami:          pulumi.String("ami-1234"),
			InstanceType: pulumi.String("c4.xlarge"),
			SpotPrice:    pulumi.String("0.03"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("CheapWorker"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}


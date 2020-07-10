package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewDefaultSubnet(ctx, "defaultAz1", &ec2.DefaultSubnetArgs{
			AvailabilityZone: pulumi.String("us-west-2a"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Default subnet for us-west-2a"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}


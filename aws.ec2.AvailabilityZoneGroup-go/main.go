package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewAvailabilityZoneGroup(ctx, "example", &ec2.AvailabilityZoneGroupArgs{
			GroupName:   pulumi.String("us-west-2-lax-1"),
			OptInStatus: pulumi.String("opted-in"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}


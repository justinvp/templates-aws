package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewCapacityReservation(ctx, "_default", &ec2.CapacityReservationArgs{
			AvailabilityZone: pulumi.String("eu-west-1a"),
			InstanceCount:    pulumi.Int(1),
			InstancePlatform: pulumi.String("Linux/UNIX"),
			InstanceType:     pulumi.String("t2.micro"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}


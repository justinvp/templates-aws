package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		web, err := ec2.NewInstance(ctx, "web", &ec2.InstanceArgs{
			Ami:              pulumi.String("ami-21f78e11"),
			AvailabilityZone: pulumi.String("us-west-2a"),
			InstanceType:     pulumi.String("t1.micro"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("HelloWorld"),
			},
		})
		if err != nil {
			return err
		}
		example, err := ec2.NewEip(ctx, "example", &ec2.EipArgs{
			Vpc: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewEipAssociation(ctx, "eipAssoc", &ec2.EipAssociationArgs{
			AllocationId: example.ID(),
			InstanceId:   web.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}


package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewEip(ctx, "lb", &ec2.EipArgs{
			Instance: pulumi.Any(aws_instance.Web.Id),
			Vpc:      pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		return nil
	})
}


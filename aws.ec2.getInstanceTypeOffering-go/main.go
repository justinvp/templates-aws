package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.GetInstanceTypeOffering(ctx, &ec2.GetInstanceTypeOfferingArgs{
			Filters: []ec2.GetInstanceTypeOfferingFilter{
				ec2.GetInstanceTypeOfferingFilter{
					Name: "instance-type",
					Values: []string{
						"t1.micro",
						"t2.micro",
						"t3.micro",
					},
				},
			},
			PreferredInstanceTypes: []string{
				"t3.micro",
				"t2.micro",
				"t1.micro",
			},
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}


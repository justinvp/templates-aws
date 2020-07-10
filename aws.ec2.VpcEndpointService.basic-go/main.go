package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewVpcEndpointService(ctx, "example", &ec2.VpcEndpointServiceArgs{
			AcceptanceRequired: pulumi.Bool(false),
			NetworkLoadBalancerArns: pulumi.StringArray{
				pulumi.Any(aws_lb.Example.Arn),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}


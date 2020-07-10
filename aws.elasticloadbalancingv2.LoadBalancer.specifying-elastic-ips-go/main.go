package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lb.NewLoadBalancer(ctx, "example", &lb.LoadBalancerArgs{
			LoadBalancerType: pulumi.String("network"),
			SubnetMappings: lb.LoadBalancerSubnetMappingArray{
				&lb.LoadBalancerSubnetMappingArgs{
					AllocationId: pulumi.Any(aws_eip.Example1.Id),
					SubnetId:     pulumi.Any(aws_subnet.Example1.Id),
				},
				&lb.LoadBalancerSubnetMappingArgs{
					AllocationId: pulumi.Any(aws_eip.Example2.Id),
					SubnetId:     pulumi.Any(aws_subnet.Example2.Id),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}


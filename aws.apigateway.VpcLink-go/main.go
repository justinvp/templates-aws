package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigateway"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleLoadBalancer, err := lb.NewLoadBalancer(ctx, "exampleLoadBalancer", &lb.LoadBalancerArgs{
			Internal:         pulumi.Bool(true),
			LoadBalancerType: pulumi.String("network"),
			SubnetMappings: lb.LoadBalancerSubnetMappingArray{
				&lb.LoadBalancerSubnetMappingArgs{
					SubnetId: pulumi.String("12345"),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = apigateway.NewVpcLink(ctx, "exampleVpcLink", &apigateway.VpcLinkArgs{
			Description: pulumi.String("example description"),
			TargetArn:   exampleLoadBalancer.Arn,
		})
		if err != nil {
			return err
		}
		return nil
	})
}


package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elb"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elb.LookupLoadBalancer(ctx, &elb.LookupLoadBalancerArgs{
			Name: lbName,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}


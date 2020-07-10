package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := lbArn
		opt1 := lbName
		_, err := lb.LookupLoadBalancer(ctx, &lb.LookupLoadBalancerArgs{
			Arn:  &opt0,
			Name: &opt1,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}


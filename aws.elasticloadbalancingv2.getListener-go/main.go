package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := listenerArn
		_, err := lb.LookupListener(ctx, &lb.LookupListenerArgs{
			Arn: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		opt1 := "default-public"
		selected, err := lb.LookupLoadBalancer(ctx, &lb.LookupLoadBalancerArgs{
			Name: &opt1,
		}, nil)
		if err != nil {
			return err
		}
		opt2 := selected.Arn
		opt3 := 443
		_, err = lb.LookupListener(ctx, &lb.LookupListenerArgs{
			LoadBalancerArn: &opt2,
			Port:            &opt3,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}


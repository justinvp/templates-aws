package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elb"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		lb, err := elb.NewLoadBalancer(ctx, "lb", &elb.LoadBalancerArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-east-1a"),
			},
			Listeners: elb.LoadBalancerListenerArray{
				&elb.LoadBalancerListenerArgs{
					InstancePort:     pulumi.Int(8000),
					InstanceProtocol: pulumi.String("http"),
					LbPort:           pulumi.Int(80),
					LbProtocol:       pulumi.String("http"),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = elb.NewLoadBalancerCookieStickinessPolicy(ctx, "foo", &elb.LoadBalancerCookieStickinessPolicyArgs{
			CookieExpirationPeriod: pulumi.Int(600),
			LbPort:                 pulumi.Int(80),
			LoadBalancer:           lb.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}


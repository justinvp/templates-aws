package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
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
					InstancePort:     pulumi.Int(25),
					InstanceProtocol: pulumi.String("tcp"),
					LbPort:           pulumi.Int(25),
					LbProtocol:       pulumi.String("tcp"),
				},
				&elb.LoadBalancerListenerArgs{
					InstancePort:     pulumi.Int(587),
					InstanceProtocol: pulumi.String("tcp"),
					LbPort:           pulumi.Int(587),
					LbProtocol:       pulumi.String("tcp"),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewProxyProtocolPolicy(ctx, "smtp", &ec2.ProxyProtocolPolicyArgs{
			InstancePorts: pulumi.StringArray{
				pulumi.String("25"),
				pulumi.String("587"),
			},
			LoadBalancer: lb.Name,
		})
		if err != nil {
			return err
		}
		return nil
	})
}


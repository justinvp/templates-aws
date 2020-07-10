package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elb"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		main, err := elb.NewLoadBalancer(ctx, "main", &elb.LoadBalancerArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-east-1c"),
			},
			Listeners: elb.LoadBalancerListenerArray{
				&elb.LoadBalancerListenerArgs{
					InstancePort:     pulumi.Int(80),
					InstanceProtocol: pulumi.String("http"),
					LbPort:           pulumi.Int(80),
					LbProtocol:       pulumi.String("http"),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = route53.NewRecord(ctx, "www", &route53.RecordArgs{
			Aliases: route53.RecordAliasArray{
				&route53.RecordAliasArgs{
					EvaluateTargetHealth: pulumi.Bool(true),
					Name:                 main.DnsName,
					ZoneId:               main.ZoneId,
				},
			},
			Name:   pulumi.String("example.com"),
			Type:   pulumi.String("A"),
			ZoneId: pulumi.Any(aws_route53_zone.Primary.Zone_id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}


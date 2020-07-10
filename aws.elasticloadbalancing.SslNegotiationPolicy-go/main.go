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
					InstanceProtocol: pulumi.String("https"),
					LbPort:           pulumi.Int(443),
					LbProtocol:       pulumi.String("https"),
					SslCertificateId: pulumi.String("arn:aws:iam::123456789012:server-certificate/certName"),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = elb.NewSslNegotiationPolicy(ctx, "foo", &elb.SslNegotiationPolicyArgs{
			Attributes: elb.SslNegotiationPolicyAttributeArray{
				&elb.SslNegotiationPolicyAttributeArgs{
					Name:  pulumi.String("Protocol-TLSv1"),
					Value: pulumi.String("false"),
				},
				&elb.SslNegotiationPolicyAttributeArgs{
					Name:  pulumi.String("Protocol-TLSv1.1"),
					Value: pulumi.String("false"),
				},
				&elb.SslNegotiationPolicyAttributeArgs{
					Name:  pulumi.String("Protocol-TLSv1.2"),
					Value: pulumi.String("true"),
				},
				&elb.SslNegotiationPolicyAttributeArgs{
					Name:  pulumi.String("Server-Defined-Cipher-Order"),
					Value: pulumi.String("true"),
				},
				&elb.SslNegotiationPolicyAttributeArgs{
					Name:  pulumi.String("ECDHE-RSA-AES128-GCM-SHA256"),
					Value: pulumi.String("true"),
				},
				&elb.SslNegotiationPolicyAttributeArgs{
					Name:  pulumi.String("AES128-GCM-SHA256"),
					Value: pulumi.String("true"),
				},
				&elb.SslNegotiationPolicyAttributeArgs{
					Name:  pulumi.String("EDH-RSA-DES-CBC3-SHA"),
					Value: pulumi.String("false"),
				},
			},
			LbPort:       pulumi.Int(443),
			LoadBalancer: lb.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}


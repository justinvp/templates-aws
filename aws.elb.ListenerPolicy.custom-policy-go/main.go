package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elb"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elb.NewLoadBalancer(ctx, "wu_tang", &elb.LoadBalancerArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-east-1a"),
			},
			Listeners: elb.LoadBalancerListenerArray{
				&elb.LoadBalancerListenerArgs{
					InstancePort:     pulumi.Int(443),
					InstanceProtocol: pulumi.String("http"),
					LbPort:           pulumi.Int(443),
					LbProtocol:       pulumi.String("https"),
					SslCertificateId: pulumi.String("arn:aws:iam::000000000000:server-certificate/wu-tang.net"),
				},
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("wu-tang"),
			},
		})
		if err != nil {
			return err
		}
		_, err = elb.NewLoadBalancerPolicy(ctx, "wu_tang_ssl", &elb.LoadBalancerPolicyArgs{
			LoadBalancerName: wu_tang.Name,
			PolicyAttributes: elb.LoadBalancerPolicyPolicyAttributeArray{
				&elb.LoadBalancerPolicyPolicyAttributeArgs{
					Name:  pulumi.String("ECDHE-ECDSA-AES128-GCM-SHA256"),
					Value: pulumi.String("true"),
				},
				&elb.LoadBalancerPolicyPolicyAttributeArgs{
					Name:  pulumi.String("Protocol-TLSv1.2"),
					Value: pulumi.String("true"),
				},
			},
			PolicyName:     pulumi.String("wu-tang-ssl"),
			PolicyTypeName: pulumi.String("SSLNegotiationPolicyType"),
		})
		if err != nil {
			return err
		}
		_, err = elb.NewListenerPolicy(ctx, "wu_tang_listener_policies_443", &elb.ListenerPolicyArgs{
			LoadBalancerName: wu_tang.Name,
			LoadBalancerPort: pulumi.Int(443),
			PolicyNames: pulumi.StringArray{
				wu_tang_ssl.PolicyName,
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}


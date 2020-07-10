package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elb"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := true
		opt1 := "my-domain.org"
		my_domain, err := iam.LookupServerCertificate(ctx, &iam.LookupServerCertificateArgs{
			Latest:     &opt0,
			NamePrefix: &opt1,
		}, nil)
		if err != nil {
			return err
		}
		_, err = elb.NewLoadBalancer(ctx, "elb", &elb.LoadBalancerArgs{
			Listeners: elb.LoadBalancerListenerArray{
				&elb.LoadBalancerListenerArgs{
					InstancePort:     pulumi.Int(8000),
					InstanceProtocol: pulumi.String("https"),
					LbPort:           pulumi.Int(443),
					LbProtocol:       pulumi.String("https"),
					SslCertificateId: pulumi.String(my_domain.Arn),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}


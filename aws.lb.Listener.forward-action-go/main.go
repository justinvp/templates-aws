package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		frontEndLoadBalancer, err := lb.NewLoadBalancer(ctx, "frontEndLoadBalancer", nil)
		if err != nil {
			return err
		}
		frontEndTargetGroup, err := lb.NewTargetGroup(ctx, "frontEndTargetGroup", nil)
		if err != nil {
			return err
		}
		_, err = lb.NewListener(ctx, "frontEndListener", &lb.ListenerArgs{
			CertificateArn: pulumi.String("arn:aws:iam::187416307283:server-certificate/test_cert_rab3wuqwgja25ct3n4jdj2tzu4"),
			DefaultActions: lb.ListenerDefaultActionArray{
				&lb.ListenerDefaultActionArgs{
					TargetGroupArn: frontEndTargetGroup.Arn,
					Type:           pulumi.String("forward"),
				},
			},
			LoadBalancerArn: frontEndLoadBalancer.Arn,
			Port:            pulumi.Int(443),
			Protocol:        pulumi.String("HTTPS"),
			SslPolicy:       pulumi.String("ELBSecurityPolicy-2016-08"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}


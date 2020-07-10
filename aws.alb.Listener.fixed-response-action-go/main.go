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
		_, err = lb.NewListener(ctx, "frontEndListener", &lb.ListenerArgs{
			DefaultActions: lb.ListenerDefaultActionArray{
				&lb.ListenerDefaultActionArgs{
					FixedResponse: &lb.ListenerDefaultActionFixedResponseArgs{
						ContentType: pulumi.String("text/plain"),
						MessageBody: pulumi.String("Fixed response content"),
						StatusCode:  pulumi.String("200"),
					},
					Type: pulumi.String("fixed-response"),
				},
			},
			LoadBalancerArn: frontEndLoadBalancer.Arn,
			Port:            pulumi.Int(80),
			Protocol:        pulumi.String("HTTP"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}


package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cognito"
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
		pool, err := cognito.NewUserPool(ctx, "pool", nil)
		if err != nil {
			return err
		}
		client, err := cognito.NewUserPoolClient(ctx, "client", nil)
		if err != nil {
			return err
		}
		domain, err := cognito.NewUserPoolDomain(ctx, "domain", nil)
		if err != nil {
			return err
		}
		_, err = lb.NewListener(ctx, "frontEndListener", &lb.ListenerArgs{
			DefaultActions: lb.ListenerDefaultActionArray{
				&lb.ListenerDefaultActionArgs{
					AuthenticateCognito: &lb.ListenerDefaultActionAuthenticateCognitoArgs{
						UserPoolArn:      pool.Arn,
						UserPoolClientId: client.ID(),
						UserPoolDomain:   domain.Domain,
					},
					Type: pulumi.String("authenticate-cognito"),
				},
				&lb.ListenerDefaultActionArgs{
					TargetGroupArn: frontEndTargetGroup.Arn,
					Type:           pulumi.String("forward"),
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


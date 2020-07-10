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
			DefaultActions: lb.ListenerDefaultActionArray{
				&lb.ListenerDefaultActionArgs{
					AuthenticateOidc: &lb.ListenerDefaultActionAuthenticateOidcArgs{
						AuthorizationEndpoint: pulumi.String("https://example.com/authorization_endpoint"),
						ClientId:              pulumi.String("client_id"),
						ClientSecret:          pulumi.String("client_secret"),
						Issuer:                pulumi.String("https://example.com"),
						TokenEndpoint:         pulumi.String("https://example.com/token_endpoint"),
						UserInfoEndpoint:      pulumi.String("https://example.com/user_info_endpoint"),
					},
					Type: pulumi.String("authenticate-oidc"),
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


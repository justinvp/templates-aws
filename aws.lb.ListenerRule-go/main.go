package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cognito"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lb.NewLoadBalancer(ctx, "frontEndLoadBalancer", nil)
		if err != nil {
			return err
		}
		frontEndListener, err := lb.NewListener(ctx, "frontEndListener", nil)
		if err != nil {
			return err
		}
		_, err = lb.NewListenerRule(ctx, "static", &lb.ListenerRuleArgs{
			Actions: lb.ListenerRuleActionArray{
				&lb.ListenerRuleActionArgs{
					TargetGroupArn: pulumi.Any(aws_lb_target_group.Static.Arn),
					Type:           pulumi.String("forward"),
				},
			},
			Conditions: lb.ListenerRuleConditionArray{
				&lb.ListenerRuleConditionArgs{
					PathPattern: &lb.ListenerRuleConditionPathPatternArgs{
						Values: pulumi.StringArray{
							pulumi.String("/static/*"),
						},
					},
				},
				&lb.ListenerRuleConditionArgs{
					HostHeader: &lb.ListenerRuleConditionHostHeaderArgs{
						Values: pulumi.StringArray{
							pulumi.String("example.com"),
						},
					},
				},
			},
			ListenerArn: frontEndListener.Arn,
			Priority:    pulumi.Int(100),
		})
		if err != nil {
			return err
		}
		_, err = lb.NewListenerRule(ctx, "hostBasedRouting", &lb.ListenerRuleArgs{
			Actions: lb.ListenerRuleActionArray{
				&lb.ListenerRuleActionArgs{
					Forward: &lb.ListenerRuleActionForwardArgs{
						Stickiness: &lb.ListenerRuleActionForwardStickinessArgs{
							Duration: pulumi.Int(600),
							Enabled:  pulumi.Bool(true),
						},
						TargetGroup: pulumi.MapArray{
							pulumi.Map{
								"arn":    pulumi.Any(aws_lb_target_group.Main.Arn),
								"weight": pulumi.Float64(80),
							},
							pulumi.Map{
								"arn":    pulumi.Any(aws_lb_target_group.Canary.Arn),
								"weight": pulumi.Float64(20),
							},
						},
					},
					Type: pulumi.String("forward"),
				},
			},
			Conditions: lb.ListenerRuleConditionArray{
				&lb.ListenerRuleConditionArgs{
					HostHeader: &lb.ListenerRuleConditionHostHeaderArgs{
						Values: pulumi.StringArray{
							pulumi.String("my-service.*.mycompany.io"),
						},
					},
				},
			},
			ListenerArn: frontEndListener.Arn,
			Priority:    pulumi.Int(99),
		})
		if err != nil {
			return err
		}
		_, err = lb.NewListenerRule(ctx, "hostBasedWeightedRouting", &lb.ListenerRuleArgs{
			Actions: lb.ListenerRuleActionArray{
				&lb.ListenerRuleActionArgs{
					TargetGroupArn: pulumi.Any(aws_lb_target_group.Static.Arn),
					Type:           pulumi.String("forward"),
				},
			},
			Conditions: lb.ListenerRuleConditionArray{
				&lb.ListenerRuleConditionArgs{
					HostHeader: &lb.ListenerRuleConditionHostHeaderArgs{
						Values: pulumi.StringArray{
							pulumi.String("my-service.*.mydomain.io"),
						},
					},
				},
			},
			ListenerArn: frontEndListener.Arn,
			Priority:    pulumi.Int(99),
		})
		if err != nil {
			return err
		}
		_, err = lb.NewListenerRule(ctx, "redirectHttpToHttps", &lb.ListenerRuleArgs{
			Actions: lb.ListenerRuleActionArray{
				&lb.ListenerRuleActionArgs{
					Redirect: &lb.ListenerRuleActionRedirectArgs{
						Port:       pulumi.String("443"),
						Protocol:   pulumi.String("HTTPS"),
						StatusCode: pulumi.String("HTTP_301"),
					},
					Type: pulumi.String("redirect"),
				},
			},
			Conditions: lb.ListenerRuleConditionArray{
				&lb.ListenerRuleConditionArgs{
					HttpHeader: &lb.ListenerRuleConditionHttpHeaderArgs{
						HttpHeaderName: pulumi.String("X-Forwarded-For"),
						Values: pulumi.StringArray{
							pulumi.String("192.168.1.*"),
						},
					},
				},
			},
			ListenerArn: frontEndListener.Arn,
		})
		if err != nil {
			return err
		}
		_, err = lb.NewListenerRule(ctx, "healthCheck", &lb.ListenerRuleArgs{
			Actions: lb.ListenerRuleActionArray{
				&lb.ListenerRuleActionArgs{
					FixedResponse: &lb.ListenerRuleActionFixedResponseArgs{
						ContentType: pulumi.String("text/plain"),
						MessageBody: pulumi.String("HEALTHY"),
						StatusCode:  pulumi.String("200"),
					},
					Type: pulumi.String("fixed-response"),
				},
			},
			Conditions: lb.ListenerRuleConditionArray{
				&lb.ListenerRuleConditionArgs{
					QueryString: pulumi.Array{
						pulumi.StringMap{
							"key":   pulumi.String("health"),
							"value": pulumi.String("check"),
						},
						pulumi.StringMap{
							"value": pulumi.String("bar"),
						},
					},
				},
			},
			ListenerArn: frontEndListener.Arn,
		})
		if err != nil {
			return err
		}
		_, err = cognito.NewUserPool(ctx, "pool", nil)
		if err != nil {
			return err
		}
		_, err = cognito.NewUserPoolClient(ctx, "client", nil)
		if err != nil {
			return err
		}
		_, err = cognito.NewUserPoolDomain(ctx, "domain", nil)
		if err != nil {
			return err
		}
		_, err = lb.NewListenerRule(ctx, "admin", &lb.ListenerRuleArgs{
			Actions: lb.ListenerRuleActionArray{
				&lb.ListenerRuleActionArgs{
					AuthenticateOidc: &lb.ListenerRuleActionAuthenticateOidcArgs{
						AuthorizationEndpoint: pulumi.String("https://example.com/authorization_endpoint"),
						ClientId:              pulumi.String("client_id"),
						ClientSecret:          pulumi.String("client_secret"),
						Issuer:                pulumi.String("https://example.com"),
						TokenEndpoint:         pulumi.String("https://example.com/token_endpoint"),
						UserInfoEndpoint:      pulumi.String("https://example.com/user_info_endpoint"),
					},
					Type: pulumi.String("authenticate-oidc"),
				},
				&lb.ListenerRuleActionArgs{
					TargetGroupArn: pulumi.Any(aws_lb_target_group.Static.Arn),
					Type:           pulumi.String("forward"),
				},
			},
			ListenerArn: frontEndListener.Arn,
		})
		if err != nil {
			return err
		}
		return nil
	})
}


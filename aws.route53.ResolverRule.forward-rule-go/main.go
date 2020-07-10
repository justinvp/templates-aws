package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := route53.NewResolverRule(ctx, "fwd", &route53.ResolverRuleArgs{
			DomainName:         pulumi.String("example.com"),
			ResolverEndpointId: pulumi.Any(aws_route53_resolver_endpoint.Foo.Id),
			RuleType:           pulumi.String("FORWARD"),
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("Prod"),
			},
			TargetIps: route53.ResolverRuleTargetIpArray{
				&route53.ResolverRuleTargetIpArgs{
					Ip: pulumi.String("123.45.67.89"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}


package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := "subdomain.example.com"
		opt1 := "SYSTEM"
		_, err := route53.LookupResolverRule(ctx, &route53.LookupResolverRuleArgs{
			DomainName: &opt0,
			RuleType:   &opt1,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}


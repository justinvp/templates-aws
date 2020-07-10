package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := route53.NewResolverRuleAssociation(ctx, "example", &route53.ResolverRuleAssociationArgs{
			ResolverRuleId: pulumi.Any(aws_route53_resolver_rule.Sys.Id),
			VpcId:          pulumi.Any(aws_vpc.Foo.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}


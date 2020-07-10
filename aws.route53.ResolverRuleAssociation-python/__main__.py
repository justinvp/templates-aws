import pulumi
import pulumi_aws as aws

example = aws.route53.ResolverRuleAssociation("example",
    resolver_rule_id=aws_route53_resolver_rule["sys"]["id"],
    vpc_id=aws_vpc["foo"]["id"])


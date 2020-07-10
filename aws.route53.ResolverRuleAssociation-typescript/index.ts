import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.route53.ResolverRuleAssociation("example", {
    resolverRuleId: aws_route53_resolver_rule_sys.id,
    vpcId: aws_vpc_foo.id,
});


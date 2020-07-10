using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Route53.ResolverRuleAssociation("example", new Aws.Route53.ResolverRuleAssociationArgs
        {
            ResolverRuleId = aws_route53_resolver_rule.Sys.Id,
            VpcId = aws_vpc.Foo.Id,
        });
    }

}


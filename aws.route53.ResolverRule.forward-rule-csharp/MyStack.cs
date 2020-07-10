using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var fwd = new Aws.Route53.ResolverRule("fwd", new Aws.Route53.ResolverRuleArgs
        {
            DomainName = "example.com",
            ResolverEndpointId = aws_route53_resolver_endpoint.Foo.Id,
            RuleType = "FORWARD",
            Tags = 
            {
                { "Environment", "Prod" },
            },
            TargetIps = 
            {
                new Aws.Route53.Inputs.ResolverRuleTargetIpArgs
                {
                    Ip = "123.45.67.89",
                },
            },
        });
    }

}


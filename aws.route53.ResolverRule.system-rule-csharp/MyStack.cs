using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var sys = new Aws.Route53.ResolverRule("sys", new Aws.Route53.ResolverRuleArgs
        {
            DomainName = "subdomain.example.com",
            RuleType = "SYSTEM",
        });
    }

}


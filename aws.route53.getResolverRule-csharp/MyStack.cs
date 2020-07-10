using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Route53.GetResolverRule.InvokeAsync(new Aws.Route53.GetResolverRuleArgs
        {
            DomainName = "subdomain.example.com",
            RuleType = "SYSTEM",
        }));
    }

}


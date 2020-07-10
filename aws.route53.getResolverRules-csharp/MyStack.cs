using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Route53.GetResolverRules.InvokeAsync(new Aws.Route53.GetResolverRulesArgs
        {
            Tags = 
            {
                
                {
                    { "Environment", "dev" },
                },
            },
        }));
    }

}


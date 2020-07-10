using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ram.ResourceShare("example", new Aws.Ram.ResourceShareArgs
        {
            AllowExternalPrincipals = true,
            Tags = 
            {
                { "Environment", "Production" },
            },
        });
    }

}


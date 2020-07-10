using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Inspector.ResourceGroup("example", new Aws.Inspector.ResourceGroupArgs
        {
            Tags = 
            {
                { "Env", "bar" },
                { "Name", "foo" },
            },
        });
    }

}


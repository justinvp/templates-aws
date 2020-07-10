using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.ServiceDiscovery.HttpNamespace("example", new Aws.ServiceDiscovery.HttpNamespaceArgs
        {
            Description = "example",
        });
    }

}


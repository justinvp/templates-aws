using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.ServiceDiscovery.PublicDnsNamespace("example", new Aws.ServiceDiscovery.PublicDnsNamespaceArgs
        {
            Description = "example",
        });
    }

}


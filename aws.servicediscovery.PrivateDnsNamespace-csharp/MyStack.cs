using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleVpc = new Aws.Ec2.Vpc("exampleVpc", new Aws.Ec2.VpcArgs
        {
            CidrBlock = "10.0.0.0/16",
        });
        var examplePrivateDnsNamespace = new Aws.ServiceDiscovery.PrivateDnsNamespace("examplePrivateDnsNamespace", new Aws.ServiceDiscovery.PrivateDnsNamespaceArgs
        {
            Description = "example",
            Vpc = exampleVpc.Id,
        });
    }

}


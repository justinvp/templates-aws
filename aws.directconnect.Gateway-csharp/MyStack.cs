using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.DirectConnect.Gateway("example", new Aws.DirectConnect.GatewayArgs
        {
            AmazonSideAsn = "64512",
        });
    }

}


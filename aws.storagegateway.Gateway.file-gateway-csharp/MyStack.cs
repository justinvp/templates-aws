using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.StorageGateway.Gateway("example", new Aws.StorageGateway.GatewayArgs
        {
            GatewayIpAddress = "1.2.3.4",
            GatewayName = "example",
            GatewayTimezone = "GMT",
            GatewayType = "FILE_S3",
        });
    }

}


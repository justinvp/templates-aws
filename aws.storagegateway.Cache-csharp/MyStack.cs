using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.StorageGateway.Cache("example", new Aws.StorageGateway.CacheArgs
        {
            DiskId = data.Aws_storagegateway_local_disk.Example.Id,
            GatewayArn = aws_storagegateway_gateway.Example.Arn,
        });
    }

}


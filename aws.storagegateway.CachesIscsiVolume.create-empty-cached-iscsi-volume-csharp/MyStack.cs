using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.StorageGateway.CachesIscsiVolume("example", new Aws.StorageGateway.CachesIscsiVolumeArgs
        {
            GatewayArn = aws_storagegateway_cache.Example.Gateway_arn,
            NetworkInterfaceId = aws_instance.Example.Private_ip,
            TargetName = "example",
            VolumeSizeInBytes = 5368709120,
        });
        // 5 GB
    }

}


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
            SnapshotId = aws_ebs_snapshot.Example.Id,
            TargetName = "example",
            VolumeSizeInBytes = aws_ebs_snapshot.Example.Volume_size * 1024 * 1024 * 1024,
        });
    }

}


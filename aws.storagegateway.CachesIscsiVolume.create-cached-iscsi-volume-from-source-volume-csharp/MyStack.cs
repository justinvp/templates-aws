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
            SourceVolumeArn = aws_storagegateway_cached_iscsi_volume.Existing.Arn,
            TargetName = "example",
            VolumeSizeInBytes = aws_storagegateway_cached_iscsi_volume.Existing.Volume_size_in_bytes,
        });
    }

}


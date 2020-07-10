using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var test = Output.Create(Aws.StorageGateway.GetLocalDisk.InvokeAsync(new Aws.StorageGateway.GetLocalDiskArgs
        {
            DiskPath = aws_volume_attachment.Test.Device_name,
            GatewayArn = aws_storagegateway_gateway.Test.Arn,
        }));
    }

}


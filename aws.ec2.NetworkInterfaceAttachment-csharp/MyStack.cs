using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var test = new Aws.Ec2.NetworkInterfaceAttachment("test", new Aws.Ec2.NetworkInterfaceAttachmentArgs
        {
            DeviceIndex = 0,
            InstanceId = aws_instance.Test.Id,
            NetworkInterfaceId = aws_network_interface.Test.Id,
        });
    }

}


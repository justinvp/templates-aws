using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var test = new Aws.Ec2.NetworkInterface("test", new Aws.Ec2.NetworkInterfaceArgs
        {
            Attachments = 
            {
                new Aws.Ec2.Inputs.NetworkInterfaceAttachmentArgs
                {
                    DeviceIndex = 1,
                    Instance = aws_instance.Test.Id,
                },
            },
            PrivateIps = 
            {
                "10.0.0.50",
            },
            SecurityGroups = 
            {
                aws_security_group.Web.Id,
            },
            SubnetId = aws_subnet.Public_a.Id,
        });
    }

}


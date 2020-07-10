using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var ami = Output.Create(Aws.GetAmi.InvokeAsync(new Aws.GetAmiArgs
        {
            Filters = 
            {
                new Aws.Inputs.GetAmiFilterArgs
                {
                    Name = "name",
                    Values = 
                    {
                        "amzn-ami-hvm-*",
                    },
                },
            },
            MostRecent = true,
            Owners = 
            {
                "amazon",
            },
        }));
        var instance = new Aws.Ec2.Instance("instance", new Aws.Ec2.InstanceArgs
        {
            Ami = ami.Apply(ami => ami.Id),
            InstanceType = "t2.micro",
            Tags = 
            {
                { "type", "test-instance" },
            },
        });
        var sg = new Aws.Ec2.SecurityGroup("sg", new Aws.Ec2.SecurityGroupArgs
        {
            Tags = 
            {
                { "type", "test-security-group" },
            },
        });
        var sgAttachment = new Aws.Ec2.NetworkInterfaceSecurityGroupAttachment("sgAttachment", new Aws.Ec2.NetworkInterfaceSecurityGroupAttachmentArgs
        {
            NetworkInterfaceId = instance.PrimaryNetworkInterfaceId,
            SecurityGroupId = sg.Id,
        });
    }

}


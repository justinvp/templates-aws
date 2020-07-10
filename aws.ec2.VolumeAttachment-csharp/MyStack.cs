using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var web = new Aws.Ec2.Instance("web", new Aws.Ec2.InstanceArgs
        {
            Ami = "ami-21f78e11",
            AvailabilityZone = "us-west-2a",
            InstanceType = "t1.micro",
            Tags = 
            {
                { "Name", "HelloWorld" },
            },
        });
        var example = new Aws.Ebs.Volume("example", new Aws.Ebs.VolumeArgs
        {
            AvailabilityZone = "us-west-2a",
            Size = 1,
        });
        var ebsAtt = new Aws.Ec2.VolumeAttachment("ebsAtt", new Aws.Ec2.VolumeAttachmentArgs
        {
            DeviceName = "/dev/sdh",
            InstanceId = web.Id,
            VolumeId = example.Id,
        });
    }

}


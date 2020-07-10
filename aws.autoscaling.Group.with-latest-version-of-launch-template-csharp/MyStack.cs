using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var foobar = new Aws.Ec2.LaunchTemplate("foobar", new Aws.Ec2.LaunchTemplateArgs
        {
            ImageId = "ami-1a2b3c",
            InstanceType = "t2.micro",
            NamePrefix = "foobar",
        });
        var bar = new Aws.AutoScaling.Group("bar", new Aws.AutoScaling.GroupArgs
        {
            AvailabilityZones = 
            {
                "us-east-1a",
            },
            DesiredCapacity = 1,
            LaunchTemplate = new Aws.AutoScaling.Inputs.GroupLaunchTemplateArgs
            {
                Id = foobar.Id,
                Version = "$Latest",
            },
            MaxSize = 1,
            MinSize = 1,
        });
    }

}


using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var ubuntu = Output.Create(Aws.GetAmi.InvokeAsync(new Aws.GetAmiArgs
        {
            Filters = 
            {
                new Aws.Inputs.GetAmiFilterArgs
                {
                    Name = "name",
                    Values = 
                    {
                        "ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*",
                    },
                },
                new Aws.Inputs.GetAmiFilterArgs
                {
                    Name = "virtualization-type",
                    Values = 
                    {
                        "hvm",
                    },
                },
            },
            MostRecent = true,
            Owners = 
            {
                "099720109477",
            },
        }));
        var asConf = new Aws.Ec2.LaunchConfiguration("asConf", new Aws.Ec2.LaunchConfigurationArgs
        {
            ImageId = ubuntu.Apply(ubuntu => ubuntu.Id),
            InstanceType = "t2.micro",
        });
    }

}


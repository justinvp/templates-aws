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
        var web = new Aws.Ec2.Instance("web", new Aws.Ec2.InstanceArgs
        {
            Ami = ubuntu.Apply(ubuntu => ubuntu.Id),
            InstanceType = "t2.micro",
            Tags = 
            {
                { "Name", "HelloWorld" },
            },
        });
    }

}


using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var foo = Output.Create(Aws.Ec2.GetInstance.InvokeAsync(new Aws.Ec2.GetInstanceArgs
        {
            Filters = 
            {
                new Aws.Ec2.Inputs.GetInstanceFilterArgs
                {
                    Name = "image-id",
                    Values = 
                    {
                        "ami-xxxxxxxx",
                    },
                },
                new Aws.Ec2.Inputs.GetInstanceFilterArgs
                {
                    Name = "tag:Name",
                    Values = 
                    {
                        "instance-name-tag",
                    },
                },
            },
            InstanceId = "i-instanceid",
        }));
    }

}


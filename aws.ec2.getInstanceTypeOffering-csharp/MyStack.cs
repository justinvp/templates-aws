using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Ec2.GetInstanceTypeOffering.InvokeAsync(new Aws.Ec2.GetInstanceTypeOfferingArgs
        {
            Filters = 
            {
                new Aws.Ec2.Inputs.GetInstanceTypeOfferingFilterArgs
                {
                    Name = "instance-type",
                    Values = 
                    {
                        "t1.micro",
                        "t2.micro",
                        "t3.micro",
                    },
                },
            },
            PreferredInstanceTypes = 
            {
                "t3.micro",
                "t2.micro",
                "t1.micro",
            },
        }));
    }

}


using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Ec2.GetInstanceTypeOfferings.InvokeAsync(new Aws.Ec2.GetInstanceTypeOfferingsArgs
        {
            Filters = 
            {
                new Aws.Ec2.Inputs.GetInstanceTypeOfferingsFilterArgs
                {
                    Name = "instance-type",
                    Values = 
                    {
                        "t2.micro",
                        "t3.micro",
                    },
                },
                new Aws.Ec2.Inputs.GetInstanceTypeOfferingsFilterArgs
                {
                    Name = "location",
                    Values = 
                    {
                        "usw2-az4",
                    },
                },
            },
            LocationType = "availability-zone-id",
        }));
    }

}


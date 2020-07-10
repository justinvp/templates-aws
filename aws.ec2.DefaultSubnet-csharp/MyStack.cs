using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var defaultAz1 = new Aws.Ec2.DefaultSubnet("defaultAz1", new Aws.Ec2.DefaultSubnetArgs
        {
            AvailabilityZone = "us-west-2a",
            Tags = 
            {
                { "Name", "Default subnet for us-west-2a" },
            },
        });
    }

}


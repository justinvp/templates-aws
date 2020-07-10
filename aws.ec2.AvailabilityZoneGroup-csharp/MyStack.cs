using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ec2.AvailabilityZoneGroup("example", new Aws.Ec2.AvailabilityZoneGroupArgs
        {
            GroupName = "us-west-2-lax-1",
            OptInStatus = "opted-in",
        });
    }

}


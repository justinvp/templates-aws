using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var @default = new Aws.Ec2.CapacityReservation("default", new Aws.Ec2.CapacityReservationArgs
        {
            AvailabilityZone = "eu-west-1a",
            InstanceCount = 1,
            InstancePlatform = "Linux/UNIX",
            InstanceType = "t2.micro",
        });
    }

}


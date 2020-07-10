using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var foo = new Aws.Ec2.SpotFleetRequest("foo", new Aws.Ec2.SpotFleetRequestArgs
        {
            IamFleetRole = "arn:aws:iam::12345678:role/spot-fleet",
            LaunchSpecifications = 
            {
                new Aws.Ec2.Inputs.SpotFleetRequestLaunchSpecificationArgs
                {
                    Ami = "ami-d06a90b0",
                    AvailabilityZone = "us-west-2a",
                    InstanceType = "m1.small",
                    KeyName = "my-key",
                },
                new Aws.Ec2.Inputs.SpotFleetRequestLaunchSpecificationArgs
                {
                    Ami = "ami-d06a90b0",
                    AvailabilityZone = "us-west-2a",
                    InstanceType = "m5.large",
                    KeyName = "my-key",
                },
            },
            SpotPrice = "0.005",
            TargetCapacity = 2,
            ValidUntil = "2019-11-04T20:44:20Z",
        });
    }

}


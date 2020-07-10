using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        // Request a Spot fleet
        var cheapCompute = new Aws.Ec2.SpotFleetRequest("cheapCompute", new Aws.Ec2.SpotFleetRequestArgs
        {
            AllocationStrategy = "diversified",
            IamFleetRole = "arn:aws:iam::12345678:role/spot-fleet",
            LaunchSpecifications = 
            {
                new Aws.Ec2.Inputs.SpotFleetRequestLaunchSpecificationArgs
                {
                    Ami = "ami-1234",
                    IamInstanceProfileArn = aws_iam_instance_profile.Example.Arn,
                    InstanceType = "m4.10xlarge",
                    PlacementTenancy = "dedicated",
                    SpotPrice = "2.793",
                },
                new Aws.Ec2.Inputs.SpotFleetRequestLaunchSpecificationArgs
                {
                    Ami = "ami-5678",
                    AvailabilityZone = "us-west-1a",
                    IamInstanceProfileArn = aws_iam_instance_profile.Example.Arn,
                    InstanceType = "m4.4xlarge",
                    KeyName = "my-key",
                    RootBlockDevice = 
                    {
                        
                        {
                            { "volumeSize", "300" },
                            { "volumeType", "gp2" },
                        },
                    },
                    SpotPrice = "1.117",
                    SubnetId = "subnet-1234",
                    Tags = 
                    {
                        { "Name", "spot-fleet-example" },
                    },
                    WeightedCapacity = "35",
                },
            },
            SpotPrice = "0.03",
            TargetCapacity = 6,
            ValidUntil = "2019-11-04T20:44:20Z",
        });
    }

}


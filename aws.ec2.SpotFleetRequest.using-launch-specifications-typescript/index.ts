import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Request a Spot fleet
const cheapCompute = new aws.ec2.SpotFleetRequest("cheap_compute", {
    allocationStrategy: "diversified",
    iamFleetRole: "arn:aws:iam::12345678:role/spot-fleet",
    launchSpecifications: [
        {
            ami: "ami-1234",
            iamInstanceProfileArn: aws_iam_instance_profile_example.arn,
            instanceType: "m4.10xlarge",
            placementTenancy: "dedicated",
            spotPrice: "2.793",
        },
        {
            ami: "ami-5678",
            availabilityZone: "us-west-1a",
            iamInstanceProfileArn: aws_iam_instance_profile_example.arn,
            instanceType: "m4.4xlarge",
            keyName: "my-key",
            rootBlockDevices: [{
                volumeSize: 300,
                volumeType: "gp2",
            }],
            spotPrice: "1.117",
            subnetId: "subnet-1234",
            tags: {
                Name: "spot-fleet-example",
            },
            weightedCapacity: "35",
        },
    ],
    spotPrice: "0.03",
    targetCapacity: 6,
    validUntil: "2019-11-04T20:44:20Z",
});


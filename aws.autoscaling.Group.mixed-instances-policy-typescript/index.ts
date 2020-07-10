import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleLaunchTemplate = new aws.ec2.LaunchTemplate("example", {
    imageId: aws_ami_example.id,
    instanceType: "c5.large",
    namePrefix: "example",
});
const exampleGroup = new aws.autoscaling.Group("example", {
    availabilityZones: ["us-east-1a"],
    desiredCapacity: 1,
    maxSize: 1,
    minSize: 1,
    mixedInstancesPolicy: {
        launchTemplate: {
            launchTemplateSpecification: {
                launchTemplateId: exampleLaunchTemplate.id,
            },
            overrides: [
                {
                    instanceType: "c4.large",
                    weightedCapacity: "3",
                },
                {
                    instanceType: "c3.large",
                    weightedCapacity: "2",
                },
            ],
        },
    },
});


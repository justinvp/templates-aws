import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.Fleet("example", {
    launchTemplateConfig: {
        launchTemplateSpecification: {
            launchTemplateId: aws_launch_template_example.id,
            version: aws_launch_template_example.latestVersion,
        },
    },
    targetCapacitySpecification: {
        defaultTargetCapacityType: "spot",
        totalTargetCapacity: 5,
    },
});


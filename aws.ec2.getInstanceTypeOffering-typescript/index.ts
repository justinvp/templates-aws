import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.ec2.getInstanceTypeOffering({
    filters: [{
        name: "instance-type",
        values: [
            "t1.micro",
            "t2.micro",
            "t3.micro",
        ],
    }],
    preferredInstanceTypes: [
        "t3.micro",
        "t2.micro",
        "t1.micro",
    ],
}, { async: true }));


import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.ec2.getInstanceTypeOfferings({
    filters: [
        {
            name: "instance-type",
            values: [
                "t2.micro",
                "t3.micro",
            ],
        },
        {
            name: "location",
            values: ["usw2-az4"],
        },
    ],
    locationType: "availability-zone-id",
}, { async: true }));


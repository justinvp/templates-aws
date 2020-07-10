import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ebsVolume = pulumi.output(aws.ebs.getVolume({
    filters: [
        {
            name: "volume-type",
            values: ["gp2"],
        },
        {
            name: "tag:Name",
            values: ["Example"],
        },
    ],
    mostRecent: true,
}, { async: true }));


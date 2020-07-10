import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ebsVolume = pulumi.output(aws.ebs.getSnapshot({
    filters: [
        {
            name: "volume-size",
            values: ["40"],
        },
        {
            name: "tag:Name",
            values: ["Example"],
        },
    ],
    mostRecent: true,
    owners: ["self"],
}, { async: true }));


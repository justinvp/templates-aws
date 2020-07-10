import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ebsVolumes = pulumi.output(aws.ebs.getSnapshotIds({
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
    owners: ["self"],
}, { async: true }));


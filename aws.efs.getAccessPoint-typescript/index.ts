import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = pulumi.output(aws.efs.getAccessPoint({
    accessPointId: "fsap-12345678",
}, { async: true }));


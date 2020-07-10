import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bar = pulumi.output(aws.ec2.getNetworkInterface({
    id: "eni-01234567",
}, { async: true }));


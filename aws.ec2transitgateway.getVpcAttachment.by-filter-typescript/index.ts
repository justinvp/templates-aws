import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.ec2transitgateway.getVpcAttachment({
    filters: [{
        name: "vpc-id",
        values: ["vpc-12345678"],
    }],
}, { async: true }));


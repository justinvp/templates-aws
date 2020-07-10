import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.ec2transitgateway.getVpcAttachment({
    id: "tgw-attach-12345678",
}, { async: true }));


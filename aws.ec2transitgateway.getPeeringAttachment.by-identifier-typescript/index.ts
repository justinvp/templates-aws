import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const attachment = pulumi.output(aws.ec2transitgateway.getPeeringAttachment({
    id: "tgw-attach-12345678",
}, { async: true }));


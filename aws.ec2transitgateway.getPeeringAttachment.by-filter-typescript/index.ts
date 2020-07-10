import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.ec2transitgateway.getPeeringAttachment({
    filters: [{
        name: "transit-gateway-attachment-id",
        values: ["tgw-attach-12345678"],
    }],
}, { async: true }));


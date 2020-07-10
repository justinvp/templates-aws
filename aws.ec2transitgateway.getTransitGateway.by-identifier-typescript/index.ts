import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.ec2transitgateway.getTransitGateway({
    id: "tgw-12345678",
}, { async: true }));


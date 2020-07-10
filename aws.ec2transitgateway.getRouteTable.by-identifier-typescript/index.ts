import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.ec2transitgateway.getRouteTable({
    id: "tgw-rtb-12345678",
}, { async: true }));


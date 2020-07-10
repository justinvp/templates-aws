import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const coipPoolId = config.require("coipPoolId");

const selected = pulumi.output(aws.ec2.getCoipPool({
    id: coipPoolId,
}, { async: true }));


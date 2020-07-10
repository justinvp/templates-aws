import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const vpcId = config.require("vpcId");

const defaultInternetGateway = pulumi.output(aws.ec2.getInternetGateway({
    filters: [{
        name: "attachment.vpc-id",
        values: [vpcId],
    }],
}, { async: true }));


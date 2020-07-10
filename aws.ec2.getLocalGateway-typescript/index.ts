import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const localGatewayId = config.require("localGatewayId");

const selected = pulumi.output(aws.ec2.getLocalGateway({
    id: localGatewayId,
}, { async: true }));


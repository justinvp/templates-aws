import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const securityGroupId = config.require("securityGroupId");

const selected = pulumi.output(aws.ec2.getSecurityGroup({
    id: securityGroupId,
}, { async: true }));
const subnet = new aws.ec2.Subnet("subnet", {
    cidrBlock: "10.0.1.0/24",
    vpcId: selected.vpcId!,
});


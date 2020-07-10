import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const subnetId = config.require("subnetId");

const selected = pulumi.output(aws.ec2.getSubnet({
    id: subnetId,
}, { async: true }));
const subnet = new aws.ec2.SecurityGroup("subnet", {
    ingress: [{
        cidrBlocks: [selected.cidrBlock!],
        fromPort: 80,
        protocol: "tcp",
        toPort: 80,
    }],
    vpcId: selected.vpcId!,
});


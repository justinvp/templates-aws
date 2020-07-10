import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const privateS3VpcEndpoint = new aws.ec2.VpcEndpoint("private_s3", {
    serviceName: "com.amazonaws.us-west-2.s3",
    vpcId: aws_vpc_foo.id,
});
const privateS3PrefixList = privateS3VpcEndpoint.prefixListId.apply(prefixListId => aws.getPrefixList({
    prefixListId: prefixListId,
}, { async: true }));
const bar = new aws.ec2.NetworkAcl("bar", {
    vpcId: aws_vpc_foo.id,
});
const privateS3NetworkAclRule = new aws.ec2.NetworkAclRule("private_s3", {
    cidrBlock: privateS3PrefixList.apply(privateS3PrefixList => privateS3PrefixList.cidrBlocks[0]),
    egress: false,
    fromPort: 443,
    networkAclId: bar.id,
    protocol: "tcp",
    ruleAction: "allow",
    ruleNumber: 200,
    toPort: 443,
});


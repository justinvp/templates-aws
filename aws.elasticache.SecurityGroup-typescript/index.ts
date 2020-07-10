import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const barEc2SecurityGroup = new aws.ec2.SecurityGroup("bar", {});
const barSecurityGroup = new aws.elasticache.SecurityGroup("bar", {
    securityGroupNames: [barEc2SecurityGroup.name],
});


import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ami = pulumi.output(aws.getAmi({
    filters: [{
        name: "name",
        values: ["amzn-ami-hvm-*"],
    }],
    mostRecent: true,
    owners: ["amazon"],
}, { async: true }));
const instance = new aws.ec2.Instance("instance", {
    ami: ami.id,
    instanceType: "t2.micro",
    tags: {
        type: "test-instance",
    },
});
const sg = new aws.ec2.SecurityGroup("sg", {
    tags: {
        type: "test-security-group",
    },
});
const sgAttachment = new aws.ec2.NetworkInterfaceSecurityGroupAttachment("sg_attachment", {
    networkInterfaceId: instance.primaryNetworkInterfaceId,
    securityGroupId: sg.id,
});


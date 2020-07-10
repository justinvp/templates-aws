import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const web = new aws.ec2.Instance("web", {
    ami: "ami-21f78e11",
    availabilityZone: "us-west-2a",
    instanceType: "t1.micro",
    tags: {
        Name: "HelloWorld",
    },
});
const example = new aws.ec2.Eip("example", {
    vpc: true,
});
const eipAssoc = new aws.ec2.EipAssociation("eip_assoc", {
    allocationId: example.id,
    instanceId: web.id,
});


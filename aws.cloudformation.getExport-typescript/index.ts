import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const subnetId = pulumi.output(aws.cloudformation.getExport({
    name: "mySubnetIdExportName",
}, { async: true }));
const web = new aws.ec2.Instance("web", {
    ami: "ami-abb07bcb",
    instanceType: "t1.micro",
    subnetId: subnetId.value,
});


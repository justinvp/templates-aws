import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const network = pulumi.output(aws.cloudformation.getStack({
    name: "my-network-stack",
}, { async: true }));
const web = new aws.ec2.Instance("web", {
    ami: "ami-abb07bcb",
    instanceType: "t1.micro",
    subnetId: network.apply(network => network.outputs["SubnetId"]),
    tags: {
        Name: "HelloWorld",
    },
});


import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ubuntu = pulumi.output(aws.getAmi({
    filters: [
        {
            name: "name",
            values: ["ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*"],
        },
        {
            name: "virtualization-type",
            values: ["hvm"],
        },
    ],
    mostRecent: true,
    owners: ["099720109477"], // Canonical
}, { async: true }));
const asConf = new aws.ec2.LaunchConfiguration("as_conf", {
    imageId: ubuntu.id,
    instanceType: "t2.micro",
});


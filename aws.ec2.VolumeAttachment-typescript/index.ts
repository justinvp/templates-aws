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
const example = new aws.ebs.Volume("example", {
    availabilityZone: "us-west-2a",
    size: 1,
});
const ebsAtt = new aws.ec2.VolumeAttachment("ebs_att", {
    deviceName: "/dev/sdh",
    instanceId: web.id,
    volumeId: example.id,
});


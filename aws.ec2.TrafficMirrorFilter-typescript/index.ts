import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.ec2.TrafficMirrorFilter("foo", {
    description: "traffic mirror filter - example",
    networkServices: ["amazon-dns"],
});


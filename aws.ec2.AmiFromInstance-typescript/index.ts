import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.AmiFromInstance("example", {
    sourceInstanceId: "i-xxxxxxxx",
});


import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const hoge = new aws.directconnect.Connection("hoge", {
    bandwidth: "1Gbps",
    location: "EqDC2",
});


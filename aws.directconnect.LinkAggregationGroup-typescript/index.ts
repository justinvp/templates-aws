import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const hoge = new aws.directconnect.LinkAggregationGroup("hoge", {
    connectionsBandwidth: "1Gbps",
    forceDestroy: true,
    location: "EqDC2",
});


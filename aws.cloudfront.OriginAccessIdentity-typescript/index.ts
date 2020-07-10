import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const originAccessIdentity = new aws.cloudfront.OriginAccessIdentity("origin_access_identity", {
    comment: "Some comment",
});


import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.storagegateway.SmbFileShare("example", {
    authentication: "ActiveDirectory",
    gatewayArn: aws_storagegateway_gateway_example.arn,
    locationArn: aws_s3_bucket_example.arn,
    roleArn: aws_iam_role_example.arn,
});


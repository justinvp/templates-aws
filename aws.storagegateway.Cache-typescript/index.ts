import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.storagegateway.Cache("example", {
    diskId: aws_storagegateway_local_disk_example.id,
    gatewayArn: aws_storagegateway_gateway_example.arn,
});


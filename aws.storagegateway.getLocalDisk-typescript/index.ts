import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = pulumi.all([aws_volume_attachment_test.deviceName, aws_storagegateway_gateway_test.arn]).apply(([deviceName, arn]) => aws.storagegateway.getLocalDisk({
    diskPath: deviceName,
    gatewayArn: arn,
}, { async: true }));


import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.storagegateway.CachesIscsiVolume("example", {
    gatewayArn: aws_storagegateway_cache_example.gatewayArn,
    networkInterfaceId: aws_instance_example.privateIp,
    targetName: "example",
    volumeSizeInBytes: 5368709120, // 5 GB
});


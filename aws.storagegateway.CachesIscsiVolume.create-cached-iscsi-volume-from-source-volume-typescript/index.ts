import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.storagegateway.CachesIscsiVolume("example", {
    gatewayArn: aws_storagegateway_cache_example.gatewayArn,
    networkInterfaceId: aws_instance_example.privateIp,
    sourceVolumeArn: aws_storagegateway_cached_iscsi_volume_existing.arn,
    targetName: "example",
    volumeSizeInBytes: aws_storagegateway_cached_iscsi_volume_existing.volumeSizeInBytes,
});


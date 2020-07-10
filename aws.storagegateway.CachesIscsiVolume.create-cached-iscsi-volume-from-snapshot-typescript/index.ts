import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.storagegateway.CachesIscsiVolume("example", {
    gatewayArn: aws_storagegateway_cache_example.gatewayArn,
    networkInterfaceId: aws_instance_example.privateIp,
    snapshotId: aws_ebs_snapshot_example.id,
    targetName: "example",
    volumeSizeInBytes: aws_ebs_snapshot_example.volumeSize.apply(volumeSize => (((volumeSize * 1024) * 1024) * 1024)),
});


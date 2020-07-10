import pulumi
import pulumi_aws as aws

test = aws.storagegateway.get_local_disk(disk_path=aws_volume_attachment["test"]["device_name"],
    gateway_arn=aws_storagegateway_gateway["test"]["arn"])


import pulumi
import pulumi_aws as aws

example = aws.storagegateway.UploadBuffer("example",
    disk_id=data["aws_storagegateway_local_disk"]["example"]["id"],
    gateway_arn=aws_storagegateway_gateway["example"]["arn"])


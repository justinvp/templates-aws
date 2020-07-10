import pulumi
import pulumi_aws as aws

example = aws.storagegateway.CachesIscsiVolume("example",
    gateway_arn=aws_storagegateway_cache["example"]["gateway_arn"],
    network_interface_id=aws_instance["example"]["private_ip"],
    target_name="example",
    volume_size_in_bytes=5368709120)
# 5 GB


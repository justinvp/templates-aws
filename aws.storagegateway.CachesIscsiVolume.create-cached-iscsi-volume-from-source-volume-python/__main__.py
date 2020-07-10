import pulumi
import pulumi_aws as aws

example = aws.storagegateway.CachesIscsiVolume("example",
    gateway_arn=aws_storagegateway_cache["example"]["gateway_arn"],
    network_interface_id=aws_instance["example"]["private_ip"],
    source_volume_arn=aws_storagegateway_cached_iscsi_volume["existing"]["arn"],
    target_name="example",
    volume_size_in_bytes=aws_storagegateway_cached_iscsi_volume["existing"]["volume_size_in_bytes"])


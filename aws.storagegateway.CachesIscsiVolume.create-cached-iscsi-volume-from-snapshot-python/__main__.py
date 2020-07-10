import pulumi
import pulumi_aws as aws

example = aws.storagegateway.CachesIscsiVolume("example",
    gateway_arn=aws_storagegateway_cache["example"]["gateway_arn"],
    network_interface_id=aws_instance["example"]["private_ip"],
    snapshot_id=aws_ebs_snapshot["example"]["id"],
    target_name="example",
    volume_size_in_bytes=aws_ebs_snapshot["example"]["volume_size"] * 1024 * 1024 * 1024)


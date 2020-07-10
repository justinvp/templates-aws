import pulumi
import pulumi_aws as aws

filter = aws.ec2.TrafficMirrorFilter("filter",
    description="traffic mirror filter - example",
    network_services=["amazon-dns"])
target = aws.ec2.TrafficMirrorTarget("target", network_load_balancer_arn=aws_lb["lb"]["arn"])
session = aws.ec2.TrafficMirrorSession("session",
    description="traffic mirror session - example",
    network_interface_id=aws_instance["test"]["primary_network_interface_id"],
    traffic_mirror_filter_id=filter.id,
    traffic_mirror_target_id=target.id)


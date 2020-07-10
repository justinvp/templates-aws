import pulumi
import pulumi_aws as aws

nlb = aws.ec2.TrafficMirrorTarget("nlb",
    description="NLB target",
    network_load_balancer_arn=aws_lb["lb"]["arn"])
eni = aws.ec2.TrafficMirrorTarget("eni",
    description="ENI target",
    network_interface_id=aws_instance["test"]["primary_network_interface_id"])


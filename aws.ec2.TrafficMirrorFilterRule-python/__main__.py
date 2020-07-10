import pulumi
import pulumi_aws as aws

filter = aws.ec2.TrafficMirrorFilter("filter",
    description="traffic mirror filter - example",
    network_services=["amazon-dns"])
ruleout = aws.ec2.TrafficMirrorFilterRule("ruleout",
    description="test rule",
    destination_cidr_block="10.0.0.0/8",
    rule_action="accept",
    rule_number=1,
    source_cidr_block="10.0.0.0/8",
    traffic_direction="egress",
    traffic_mirror_filter_id=filter.id)
rulein = aws.ec2.TrafficMirrorFilterRule("rulein",
    description="test rule",
    destination_cidr_block="10.0.0.0/8",
    destination_port_range={
        "from_port": 22,
        "to_port": 53,
    },
    protocol=6,
    rule_action="accept",
    rule_number=1,
    source_cidr_block="10.0.0.0/8",
    source_port_range={
        "from_port": 0,
        "to_port": 10,
    },
    traffic_direction="ingress",
    traffic_mirror_filter_id=filter.id)


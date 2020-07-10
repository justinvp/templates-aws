import pulumi
import pulumi_aws as aws

bar_network_acl = aws.ec2.NetworkAcl("barNetworkAcl", vpc_id=aws_vpc["foo"]["id"])
bar_network_acl_rule = aws.ec2.NetworkAclRule("barNetworkAclRule",
    network_acl_id=bar_network_acl.id,
    rule_number=200,
    egress=False,
    protocol="tcp",
    rule_action="allow",
    cidr_block=aws_vpc["foo"]["cidr_block"],
    from_port=22,
    to_port=22)


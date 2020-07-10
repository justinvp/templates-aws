import pulumi
import pulumi_aws as aws

private_s3_vpc_endpoint = aws.ec2.VpcEndpoint("privateS3VpcEndpoint",
    service_name="com.amazonaws.us-west-2.s3",
    vpc_id=aws_vpc["foo"]["id"])
private_s3_prefix_list = private_s3_vpc_endpoint.prefix_list_id.apply(lambda prefix_list_id: aws.get_prefix_list(prefix_list_id=prefix_list_id))
bar = aws.ec2.NetworkAcl("bar", vpc_id=aws_vpc["foo"]["id"])
private_s3_network_acl_rule = aws.ec2.NetworkAclRule("privateS3NetworkAclRule",
    cidr_block=private_s3_prefix_list.cidr_blocks[0],
    egress=False,
    from_port=443,
    network_acl_id=bar.id,
    protocol="tcp",
    rule_action="allow",
    rule_number=200,
    to_port=443)


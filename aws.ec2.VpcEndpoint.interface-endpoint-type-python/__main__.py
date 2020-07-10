import pulumi
import pulumi_aws as aws

ec2 = aws.ec2.VpcEndpoint("ec2",
    private_dns_enabled=True,
    security_group_ids=[aws_security_group["sg1"]["id"]],
    service_name="com.amazonaws.us-west-2.ec2",
    vpc_endpoint_type="Interface",
    vpc_id=aws_vpc["main"]["id"])


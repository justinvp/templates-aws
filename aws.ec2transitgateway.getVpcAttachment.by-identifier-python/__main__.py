import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.get_vpc_attachment(id="tgw-attach-12345678")


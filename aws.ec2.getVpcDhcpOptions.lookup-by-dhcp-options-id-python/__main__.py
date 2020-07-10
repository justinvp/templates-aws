import pulumi
import pulumi_aws as aws

example = aws.ec2.get_vpc_dhcp_options(dhcp_options_id="dopts-12345678")


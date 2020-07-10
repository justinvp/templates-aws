import pulumi
import pulumi_aws as aws

attachment = aws.ec2transitgateway.get_peering_attachment(id="tgw-attach-12345678")


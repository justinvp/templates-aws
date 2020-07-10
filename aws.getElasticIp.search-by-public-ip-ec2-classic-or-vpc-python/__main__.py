import pulumi
import pulumi_aws as aws

by_public_ip = aws.get_elastic_ip(public_ip="1.2.3.4")


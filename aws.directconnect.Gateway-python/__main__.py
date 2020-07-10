import pulumi
import pulumi_aws as aws

example = aws.directconnect.Gateway("example", amazon_side_asn="64512")


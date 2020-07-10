import pulumi
import pulumi_aws as aws

example = aws.storagegateway.Gateway("example",
    gateway_ip_address="1.2.3.4",
    gateway_name="example",
    gateway_timezone="GMT",
    gateway_type="STORED")


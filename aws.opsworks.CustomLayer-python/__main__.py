import pulumi
import pulumi_aws as aws

custlayer = aws.opsworks.CustomLayer("custlayer",
    short_name="awesome",
    stack_id=aws_opsworks_stack["main"]["id"])


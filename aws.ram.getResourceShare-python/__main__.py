import pulumi
import pulumi_aws as aws

example = aws.ram.get_resource_share(name="example",
    resource_owner="SELF")


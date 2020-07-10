import pulumi
import pulumi_aws as aws

example = aws.ram.ResourceAssociation("example",
    resource_arn=aws_subnet["example"]["arn"],
    resource_share_arn=aws_ram_resource_share["example"]["arn"])


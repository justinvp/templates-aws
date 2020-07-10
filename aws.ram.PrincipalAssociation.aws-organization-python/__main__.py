import pulumi
import pulumi_aws as aws

example = aws.ram.PrincipalAssociation("example",
    principal=aws_organizations_organization["example"]["arn"],
    resource_share_arn=aws_ram_resource_share["example"]["arn"])


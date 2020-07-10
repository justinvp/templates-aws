import pulumi
import pulumi_aws as aws

contractors = aws.workspaces.IpGroup("contractors", description="Contractors IP access control group")


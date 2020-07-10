import pulumi
import pulumi_aws as aws

example = aws.workspaces.get_bundle(name="Value with Windows 10 and Office 2016",
    owner="AMAZON")


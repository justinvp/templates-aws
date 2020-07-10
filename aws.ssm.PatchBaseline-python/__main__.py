import pulumi
import pulumi_aws as aws

production = aws.ssm.PatchBaseline("production", approved_patches=["KB123456"])


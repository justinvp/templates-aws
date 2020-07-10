import pulumi
import pulumi_aws as aws

sfn_activity = aws.sfn.Activity("sfnActivity")


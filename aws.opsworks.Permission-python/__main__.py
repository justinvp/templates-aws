import pulumi
import pulumi_aws as aws

my_stack_permission = aws.opsworks.Permission("myStackPermission",
    allow_ssh=True,
    allow_sudo=True,
    level="iam_only",
    stack_id=aws_opsworks_stack["stack"]["id"],
    user_arn=aws_iam_user["user"]["arn"])


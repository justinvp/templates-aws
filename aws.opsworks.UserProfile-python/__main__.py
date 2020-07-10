import pulumi
import pulumi_aws as aws

my_profile = aws.opsworks.UserProfile("myProfile",
    ssh_username="my_user",
    user_arn=aws_iam_user["user"]["arn"])


import pulumi
import pulumi_aws as aws

example = aws.iam.get_user(user_name="an_example_user_name")


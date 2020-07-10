import pulumi
import pulumi_aws as aws

example = aws.quicksight.User("example",
    email="author@example.com",
    identity_type="IAM",
    user_name="an-author",
    user_role="AUTHOR")


import pulumi
import pulumi_aws as aws

example = aws.codebuild.SourceCredential("example",
    auth_type="BASIC_AUTH",
    server_type="BITBUCKET",
    token="example",
    user_name="test-user")


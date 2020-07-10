import pulumi
import pulumi_aws as aws

example = aws.codebuild.SourceCredential("example",
    auth_type="PERSONAL_ACCESS_TOKEN",
    server_type="GITHUB",
    token="example")


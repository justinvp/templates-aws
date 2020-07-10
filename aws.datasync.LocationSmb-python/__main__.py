import pulumi
import pulumi_aws as aws

example = aws.datasync.LocationSmb("example",
    agent_arns=[aws_datasync_agent["example"]["arn"]],
    password="ANotGreatPassword",
    server_hostname="smb.example.com",
    subdirectory="/exported/path",
    user="Guest")


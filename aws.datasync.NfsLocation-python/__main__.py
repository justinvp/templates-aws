import pulumi
import pulumi_aws as aws

example = aws.datasync.NfsLocation("example",
    on_prem_config={
        "agent_arns": [aws_datasync_agent["example"]["arn"]],
    },
    server_hostname="nfs.example.com",
    subdirectory="/exported/path")


import pulumi
import pulumi_aws as aws

example = aws.codebuild.Webhook("example",
    filter_groups=[{
        "filter": [
            {
                "pattern": "PUSH",
                "type": "EVENT",
            },
            {
                "pattern": "master",
                "type": "HEAD_REF",
            },
        ],
    }],
    project_name=aws_codebuild_project["example"]["name"])


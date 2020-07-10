import pulumi
import pulumi_aws as aws

example = aws.glue.Trigger("example",
    actions=[{
        "jobName": aws_glue_job["example"]["name"],
    }],
    type="ON_DEMAND")


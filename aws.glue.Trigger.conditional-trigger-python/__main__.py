import pulumi
import pulumi_aws as aws

example = aws.glue.Trigger("example",
    actions=[{
        "jobName": aws_glue_job["example1"]["name"],
    }],
    predicate={
        "conditions": [{
            "jobName": aws_glue_job["example2"]["name"],
            "state": "SUCCEEDED",
        }],
    },
    type="CONDITIONAL")


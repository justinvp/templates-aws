import pulumi
import pulumi_aws as aws

example = aws.glue.Trigger("example",
    actions=[{
        "jobName": aws_glue_job["example1"]["name"],
    }],
    predicate={
        "conditions": [{
            "crawlState": "SUCCEEDED",
            "crawlerName": aws_glue_crawler["example2"]["name"],
        }],
    },
    type="CONDITIONAL")


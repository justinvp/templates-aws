import pulumi
import pulumi_aws as aws

example = aws.glue.Workflow("example")
example_start = aws.glue.Trigger("example-start",
    actions=[{
        "jobName": "example-job",
    }],
    type="ON_DEMAND",
    workflow_name=example.name)
example_inner = aws.glue.Trigger("example-inner",
    actions=[{
        "jobName": "another-example-job",
    }],
    predicate={
        "conditions": [{
            "jobName": "example-job",
            "state": "SUCCEEDED",
        }],
    },
    type="CONDITIONAL",
    workflow_name=example.name)


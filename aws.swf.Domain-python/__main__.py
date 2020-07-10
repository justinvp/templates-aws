import pulumi
import pulumi_aws as aws

foo = aws.swf.Domain("foo",
    description="SWF Domain",
    workflow_execution_retention_period_in_days=30)


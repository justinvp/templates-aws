import pulumi
import pulumi_aws as aws

example_log_group = aws.cloudwatch.LogGroup("exampleLogGroup", retention_in_days=14)
example_job = aws.glue.Job("exampleJob", default_arguments={
    "--continuous-log-logGroup": example_log_group.name,
    "--enable-continuous-cloudwatch-log": "true",
    "--enable-continuous-log-filter": "true",
    "--enable-metrics": "",
})


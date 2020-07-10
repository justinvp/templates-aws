import pulumi
import pulumi_aws as aws

dada = aws.cloudwatch.LogGroup("dada")
yada = aws.cloudwatch.LogMetricFilter("yada",
    log_group_name=dada.name,
    metric_transformation={
        "name": "EventCount",
        "namespace": "YourNamespace",
        "value": "1",
    },
    pattern="")


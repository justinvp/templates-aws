import pulumi
import pulumi_aws as aws

example = aws.neptune.ClusterParameterGroup("example",
    description="neptune cluster parameter group",
    family="neptune1",
    parameters=[{
        "name": "neptune_enable_audit_log",
        "value": 1,
    }])


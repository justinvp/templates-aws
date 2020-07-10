import pulumi
import pulumi_aws as aws

example = aws.docdb.ClusterParameterGroup("example",
    description="docdb cluster parameter group",
    family="docdb3.6",
    parameters=[{
        "name": "tls",
        "value": "enabled",
    }])


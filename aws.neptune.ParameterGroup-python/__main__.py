import pulumi
import pulumi_aws as aws

example = aws.neptune.ParameterGroup("example",
    family="neptune1",
    parameters=[{
        "name": "neptune_query_timeout",
        "value": "25",
    }])


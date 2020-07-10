import pulumi
import pulumi_aws as aws

default = aws.elasticache.ParameterGroup("default",
    family="redis2.8",
    parameters=[
        {
            "name": "activerehashing",
            "value": "yes",
        },
        {
            "name": "min-slaves-to-write",
            "value": "2",
        },
    ])


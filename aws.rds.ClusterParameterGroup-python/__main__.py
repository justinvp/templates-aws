import pulumi
import pulumi_aws as aws

default = aws.rds.ClusterParameterGroup("default",
    description="RDS default cluster parameter group",
    family="aurora5.6",
    parameters=[
        {
            "name": "character_set_server",
            "value": "utf8",
        },
        {
            "name": "character_set_client",
            "value": "utf8",
        },
    ])


import pulumi
import pulumi_aws as aws

example = aws.dax.ParameterGroup("example", parameters=[
    {
        "name": "query-ttl-millis",
        "value": "100000",
    },
    {
        "name": "record-ttl-millis",
        "value": "100000",
    },
])


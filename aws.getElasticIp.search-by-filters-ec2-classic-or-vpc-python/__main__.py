import pulumi
import pulumi_aws as aws

by_filter = aws.get_elastic_ip(filters=[{
    "name": "tag:Name",
    "values": ["exampleNameTagValue"],
}])


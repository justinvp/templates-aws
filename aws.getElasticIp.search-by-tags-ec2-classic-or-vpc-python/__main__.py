import pulumi
import pulumi_aws as aws

by_tags = aws.get_elastic_ip(tags={
    "Name": "exampleNameTagValue",
})


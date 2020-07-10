import pulumi
import pulumi_aws as aws

example = aws.elasticsearch.Domain("example",
    cluster_config={
        "cluster_config": "r4.large.elasticsearch",
    },
    elasticsearch_version="1.5",
    snapshot_options={
        "snapshot_options": 23,
    },
    tags={
        "Domain": "TestDomain",
    })


import pulumi
import pulumi_aws as aws

hoge = aws.directconnect.LinkAggregationGroup("hoge",
    connections_bandwidth="1Gbps",
    force_destroy=True,
    location="EqDC2")


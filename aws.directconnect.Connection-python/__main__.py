import pulumi
import pulumi_aws as aws

hoge = aws.directconnect.Connection("hoge",
    bandwidth="1Gbps",
    location="EqDC2")


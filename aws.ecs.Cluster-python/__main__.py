import pulumi
import pulumi_aws as aws

foo = aws.ecs.Cluster("foo")


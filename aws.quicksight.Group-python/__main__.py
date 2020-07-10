import pulumi
import pulumi_aws as aws

example = aws.quicksight.Group("example", group_name="tf-example")


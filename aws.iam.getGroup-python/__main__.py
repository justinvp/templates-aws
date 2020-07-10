import pulumi
import pulumi_aws as aws

example = aws.iam.get_group(group_name="an_example_group_name")


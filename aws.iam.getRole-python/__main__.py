import pulumi
import pulumi_aws as aws

example = aws.iam.get_role(name="an_example_role_name")


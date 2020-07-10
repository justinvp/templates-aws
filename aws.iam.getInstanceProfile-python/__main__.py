import pulumi
import pulumi_aws as aws

example = aws.iam.get_instance_profile(name="an_example_instance_profile_name")


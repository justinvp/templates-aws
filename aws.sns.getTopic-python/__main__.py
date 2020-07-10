import pulumi
import pulumi_aws as aws

example = aws.sns.get_topic(name="an_example_topic")


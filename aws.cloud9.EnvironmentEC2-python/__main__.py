import pulumi
import pulumi_aws as aws

example = aws.cloud9.EnvironmentEC2("example", instance_type="t2.micro")


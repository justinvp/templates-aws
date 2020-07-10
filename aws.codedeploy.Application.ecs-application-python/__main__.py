import pulumi
import pulumi_aws as aws

example = aws.codedeploy.Application("example", compute_platform="ECS")


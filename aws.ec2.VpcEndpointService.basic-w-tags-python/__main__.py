import pulumi
import pulumi_aws as aws

example = aws.ec2.VpcEndpointService("example",
    acceptance_required=False,
    network_load_balancer_arns=[aws_lb["example"]["arn"]],
    tags={
        "Environment": "test",
    })


import pulumi
import pulumi_aws as aws

example = aws.globalaccelerator.EndpointGroup("example",
    endpoint_configurations=[{
        "endpoint_id": aws_lb["example"]["arn"],
        "weight": 100,
    }],
    listener_arn=aws_globalaccelerator_listener["example"]["id"])


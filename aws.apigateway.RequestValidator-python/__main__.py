import pulumi
import pulumi_aws as aws

example = aws.apigateway.RequestValidator("example",
    rest_api=aws_api_gateway_rest_api["example"]["id"],
    validate_request_body=True,
    validate_request_parameters=True)


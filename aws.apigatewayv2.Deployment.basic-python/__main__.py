import pulumi
import pulumi_aws as aws

example = aws.apigatewayv2.Deployment("example",
    api_id=aws_apigatewayv2_route["example"]["api_id"],
    description="Example deployment")


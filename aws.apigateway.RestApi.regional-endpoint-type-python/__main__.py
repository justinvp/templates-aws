import pulumi
import pulumi_aws as aws

example = aws.apigateway.RestApi("example", endpoint_configuration={
    "types": "REGIONAL",
})


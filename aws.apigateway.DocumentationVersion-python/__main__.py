import pulumi
import pulumi_aws as aws

example_rest_api = aws.apigateway.RestApi("exampleRestApi")
example_documentation_version = aws.apigateway.DocumentationVersion("exampleDocumentationVersion",
    description="Example description",
    rest_api_id=example_rest_api.id,
    version="example_version",
    opts=ResourceOptions(depends_on=["aws_api_gateway_documentation_part.example"]))
example_documentation_part = aws.apigateway.DocumentationPart("exampleDocumentationPart",
    location={
        "type": "API",
    },
    properties="{\"description\":\"Example\"}",
    rest_api_id=example_rest_api.id)


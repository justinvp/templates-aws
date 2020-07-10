import pulumi
import pulumi_aws as aws

example_rest_api = aws.apigateway.RestApi("exampleRestApi")
example_documentation_part = aws.apigateway.DocumentationPart("exampleDocumentationPart",
    location={
        "method": "GET",
        "path": "/example",
        "type": "METHOD",
    },
    properties="{\"description\":\"Example description\"}",
    rest_api_id=example_rest_api.id)


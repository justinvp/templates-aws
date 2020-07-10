import pulumi
import pulumi_aws as aws

example = aws.apigatewayv2.Model("example",
    api_id=aws_apigatewayv2_api["example"]["id"],
    content_type="application/json",
    schema="""{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "title": "ExampleModel",
  "type": "object",
  "properties": {
    "id": { "type": "string" }
  }
}

""")


import pulumi
import pulumi_aws as aws

my_demo_api = aws.apigateway.RestApi("myDemoAPI", description="This is my API for demonstration purposes")
my_demo_model = aws.apigateway.Model("myDemoModel",
    content_type="application/json",
    description="a JSON schema",
    rest_api=my_demo_api.id,
    schema="""{
  "type": "object"
}

""")


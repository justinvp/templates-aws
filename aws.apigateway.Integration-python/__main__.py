import pulumi
import pulumi_aws as aws

my_demo_api = aws.apigateway.RestApi("myDemoAPI", description="This is my API for demonstration purposes")
my_demo_resource = aws.apigateway.Resource("myDemoResource",
    parent_id=my_demo_api.root_resource_id,
    path_part="mydemoresource",
    rest_api=my_demo_api.id)
my_demo_method = aws.apigateway.Method("myDemoMethod",
    authorization="NONE",
    http_method="GET",
    resource_id=my_demo_resource.id,
    rest_api=my_demo_api.id)
my_demo_integration = aws.apigateway.Integration("myDemoIntegration",
    cache_key_parameters=["method.request.path.param"],
    cache_namespace="foobar",
    http_method=my_demo_method.http_method,
    request_parameters={
        "integration.request.header.X-Authorization": "'static'",
    },
    request_templates={
        "application/xml": """{
   "body" : $input.json('$')
}

""",
    },
    resource_id=my_demo_resource.id,
    rest_api=my_demo_api.id,
    timeout_milliseconds=29000,
    type="MOCK")


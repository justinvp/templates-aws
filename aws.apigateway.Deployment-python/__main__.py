import pulumi
import pulumi_aws as aws

my_demo_api = aws.apigateway.RestApi("myDemoAPI", description="This is my API for demonstration purposes")
my_demo_resource = aws.apigateway.Resource("myDemoResource",
    rest_api=my_demo_api.id,
    parent_id=my_demo_api.root_resource_id,
    path_part="test")
my_demo_method = aws.apigateway.Method("myDemoMethod",
    rest_api=my_demo_api.id,
    resource_id=my_demo_resource.id,
    http_method="GET",
    authorization="NONE")
my_demo_integration = aws.apigateway.Integration("myDemoIntegration",
    rest_api=my_demo_api.id,
    resource_id=my_demo_resource.id,
    http_method=my_demo_method.http_method,
    type="MOCK")
my_demo_deployment = aws.apigateway.Deployment("myDemoDeployment",
    rest_api=my_demo_api.id,
    stage_name="test",
    variables={
        "answer": "42",
    },
    opts=ResourceOptions(depends_on=[my_demo_integration]))


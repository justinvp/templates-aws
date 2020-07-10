import pulumi
import pulumi_aws as aws

my_demo_api = aws.apigateway.RestApi("myDemoAPI", description="This is my API for demonstration purposes")
my_demo_resource = aws.apigateway.Resource("myDemoResource",
    parent_id=my_demo_api.root_resource_id,
    path_part="mydemoresource",
    rest_api=my_demo_api.id)


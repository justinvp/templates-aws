import pulumi
import pulumi_aws as aws

test_rest_api = aws.apigateway.RestApi("testRestApi", description="This is my API for demonstration purposes")
test_deployment = aws.apigateway.Deployment("testDeployment",
    rest_api=test_rest_api.id,
    stage_name="dev",
    opts=ResourceOptions(depends_on=["aws_api_gateway_integration.test"]))
test_stage = aws.apigateway.Stage("testStage",
    deployment=test_deployment.id,
    rest_api=test_rest_api.id,
    stage_name="prod")
test_resource = aws.apigateway.Resource("testResource",
    parent_id=test_rest_api.root_resource_id,
    path_part="mytestresource",
    rest_api=test_rest_api.id)
test_method = aws.apigateway.Method("testMethod",
    authorization="NONE",
    http_method="GET",
    resource_id=test_resource.id,
    rest_api=test_rest_api.id)
method_settings = aws.apigateway.MethodSettings("methodSettings",
    method_path=pulumi.Output.all(test_resource.path_part, test_method.http_method).apply(lambda path_part, http_method: f"{path_part}/{http_method}"),
    rest_api=test_rest_api.id,
    settings={
        "loggingLevel": "INFO",
        "metricsEnabled": True,
    },
    stage_name=test_stage.stage_name)
test_integration = aws.apigateway.Integration("testIntegration",
    http_method=test_method.http_method,
    resource_id=test_resource.id,
    rest_api=test_rest_api.id,
    type="MOCK")


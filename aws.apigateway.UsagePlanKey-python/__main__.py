import pulumi
import pulumi_aws as aws

test = aws.apigateway.RestApi("test")
myusageplan = aws.apigateway.UsagePlan("myusageplan", api_stages=[{
    "api_id": test.id,
    "stage": aws_api_gateway_deployment["foo"]["stage_name"],
}])
mykey = aws.apigateway.ApiKey("mykey")
main = aws.apigateway.UsagePlanKey("main",
    key_id=mykey.id,
    key_type="API_KEY",
    usage_plan_id=myusageplan.id)


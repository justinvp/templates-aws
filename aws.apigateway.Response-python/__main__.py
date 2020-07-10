import pulumi
import pulumi_aws as aws

main = aws.apigateway.RestApi("main")
test = aws.apigateway.Response("test",
    response_parameters={
        "gatewayresponse.header.Authorization": "'Basic'",
    },
    response_templates={
        "application/json": "{'message':$context.error.messageString}",
    },
    response_type="UNAUTHORIZED",
    rest_api_id=main.id,
    status_code="401")


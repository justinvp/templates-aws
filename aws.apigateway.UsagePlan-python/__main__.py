import pulumi
import pulumi_aws as aws

myapi = aws.apigateway.RestApi("myapi")
dev = aws.apigateway.Deployment("dev",
    rest_api=myapi.id,
    stage_name="dev")
prod = aws.apigateway.Deployment("prod",
    rest_api=myapi.id,
    stage_name="prod")
my_usage_plan = aws.apigateway.UsagePlan("myUsagePlan",
    api_stages=[
        {
            "api_id": myapi.id,
            "stage": dev.stage_name,
        },
        {
            "api_id": myapi.id,
            "stage": prod.stage_name,
        },
    ],
    description="my description",
    product_code="MYCODE",
    quota_settings={
        "limit": 20,
        "offset": 2,
        "period": "WEEK",
    },
    throttle_settings={
        "burstLimit": 5,
        "rate_limit": 10,
    })


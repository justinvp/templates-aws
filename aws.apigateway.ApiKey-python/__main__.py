import pulumi
import pulumi_aws as aws

my_demo_api_key = aws.apigateway.ApiKey("myDemoApiKey")


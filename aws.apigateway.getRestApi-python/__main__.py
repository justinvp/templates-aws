import pulumi
import pulumi_aws as aws

my_rest_api = aws.apigateway.get_rest_api(name="my-rest-api")


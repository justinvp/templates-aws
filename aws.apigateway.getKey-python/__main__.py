import pulumi
import pulumi_aws as aws

my_api_key = aws.apigateway.get_key(id="ru3mpjgse6")


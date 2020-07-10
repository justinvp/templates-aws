import pulumi
import pulumi_aws as aws

selected_rest_api = aws.apigateway.get_rest_api(name=var["api_gateway_name"])
selected_user_pools = aws.cognito.get_user_pools(name=var["cognito_user_pool_name"])
cognito = aws.apigateway.Authorizer("cognito",
    provider_arns=selected_user_pools.arns,
    rest_api=selected_rest_api.id,
    type="COGNITO_USER_POOLS")


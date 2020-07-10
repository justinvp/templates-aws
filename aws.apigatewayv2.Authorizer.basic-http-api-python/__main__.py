import pulumi
import pulumi_aws as aws

example = aws.apigatewayv2.Authorizer("example",
    api_id=aws_apigatewayv2_api["example"]["id"],
    authorizer_type="JWT",
    identity_sources=["$request.header.Authorization"],
    jwt_configuration={
        "audience": ["example"],
        "issuer": f"https://{aws_cognito_user_pool['example']['endpoint']}",
    })


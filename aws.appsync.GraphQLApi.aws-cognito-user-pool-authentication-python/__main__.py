import pulumi
import pulumi_aws as aws

example = aws.appsync.GraphQLApi("example",
    authentication_type="AMAZON_COGNITO_USER_POOLS",
    user_pool_config={
        "awsRegion": data["aws_region"]["current"]["name"],
        "default_action": "DENY",
        "user_pool_id": aws_cognito_user_pool["example"]["id"],
    })


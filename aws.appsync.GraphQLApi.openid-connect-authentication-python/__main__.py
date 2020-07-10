import pulumi
import pulumi_aws as aws

example = aws.appsync.GraphQLApi("example",
    authentication_type="OPENID_CONNECT",
    openid_connect_config={
        "issuer": "https://example.com",
    })


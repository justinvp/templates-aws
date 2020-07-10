import pulumi
import pulumi_aws as aws

example = aws.cognito.UserPool("example", auto_verified_attributes=["email"])
example_provider = aws.cognito.IdentityProvider("exampleProvider",
    attribute_mapping={
        "email": "email",
        "username": "sub",
    },
    provider_details={
        "authorize_scopes": "email",
        "client_id": "your client_id",
        "client_secret": "your client_secret",
    },
    provider_name="Google",
    provider_type="Google",
    user_pool_id=example.id)


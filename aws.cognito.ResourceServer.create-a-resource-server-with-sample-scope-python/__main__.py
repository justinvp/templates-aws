import pulumi
import pulumi_aws as aws

pool = aws.cognito.UserPool("pool")
resource = aws.cognito.ResourceServer("resource",
    identifier="https://example.com",
    scopes=[{
        "scopeDescription": "a Sample Scope Description",
        "scopeName": "sample-scope",
    }],
    user_pool_id=pool.id)


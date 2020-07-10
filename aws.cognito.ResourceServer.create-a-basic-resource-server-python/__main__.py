import pulumi
import pulumi_aws as aws

pool = aws.cognito.UserPool("pool")
resource = aws.cognito.ResourceServer("resource",
    identifier="https://example.com",
    user_pool_id=pool.id)


import pulumi
import pulumi_aws as aws

pool = aws.cognito.UserPool("pool")
client = aws.cognito.UserPoolClient("client", user_pool_id=pool.id)


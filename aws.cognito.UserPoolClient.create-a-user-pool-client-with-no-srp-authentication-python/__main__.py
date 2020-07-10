import pulumi
import pulumi_aws as aws

pool = aws.cognito.UserPool("pool")
client = aws.cognito.UserPoolClient("client",
    explicit_auth_flows=["ADMIN_NO_SRP_AUTH"],
    generate_secret=True,
    user_pool_id=pool.id)


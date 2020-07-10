import pulumi
import pulumi_aws as aws

example = aws.appsync.GraphQLApi("example",
    additional_authentication_providers=[{
        "authentication_type": "AWS_IAM",
    }],
    authentication_type="API_KEY")


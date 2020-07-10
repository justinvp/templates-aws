import pulumi
import pulumi_aws as aws

example = aws.appsync.GraphQLApi("example", authentication_type="API_KEY")


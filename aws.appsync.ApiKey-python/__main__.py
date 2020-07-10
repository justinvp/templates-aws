import pulumi
import pulumi_aws as aws

example_graph_ql_api = aws.appsync.GraphQLApi("exampleGraphQLApi", authentication_type="API_KEY")
example_api_key = aws.appsync.ApiKey("exampleApiKey",
    api_id=example_graph_ql_api.id,
    expires="2018-05-03T04:00:00Z")


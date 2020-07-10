import pulumi
import pulumi_aws as aws

example = aws.appsync.GraphQLApi("example",
    authentication_type="AWS_IAM",
    schema="""schema {
	query: Query
}
type Query {
  test: Int
}

""")


import pulumi
import pulumi_aws as aws

test_graph_ql_api = aws.appsync.GraphQLApi("testGraphQLApi",
    authentication_type="API_KEY",
    schema="""type Mutation {
	putPost(id: ID!, title: String!): Post
}

type Post {
	id: ID!
	title: String!
}

type Query {
	singlePost(id: ID!): Post
}

schema {
	query: Query
	mutation: Mutation
}

""")
test_data_source = aws.appsync.DataSource("testDataSource",
    api_id=test_graph_ql_api.id,
    http_config={
        "endpoint": "http://example.com",
    },
    type="HTTP")
# UNIT type resolver (default)
test_resolver = aws.appsync.Resolver("testResolver",
    api_id=test_graph_ql_api.id,
    caching_config={
        "cachingKeys": [
            "$context.identity.sub",
            "$context.arguments.id",
        ],
        "ttl": 60,
    },
    data_source=test_data_source.name,
    field="singlePost",
    request_template="""{
    "version": "2018-05-29",
    "method": "GET",
    "resourcePath": "/",
    "params":{
        "headers": $utils.http.copyheaders($ctx.request.headers)
    }
}

""",
    response_template="""#if($ctx.result.statusCode == 200)
    $ctx.result.body
#else
    $utils.appendError($ctx.result.body, $ctx.result.statusCode)
#end

""",
    type="Query")
# PIPELINE type resolver
mutation_pipeline_test = aws.appsync.Resolver("mutationPipelineTest",
    api_id=test_graph_ql_api.id,
    field="pipelineTest",
    kind="PIPELINE",
    pipeline_config={
        "functions": [
            aws_appsync_function["test1"]["function_id"],
            aws_appsync_function["test2"]["function_id"],
            aws_appsync_function["test3"]["function_id"],
        ],
    },
    request_template="{}",
    response_template="$util.toJson($ctx.result)",
    type="Mutation")


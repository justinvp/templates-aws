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
test_function = aws.appsync.Function("testFunction",
    api_id=test_graph_ql_api.id,
    data_source=test_data_source.name,
    name="tf_example",
    request_mapping_template="""{
    "version": "2018-05-29",
    "method": "GET",
    "resourcePath": "/",
    "params":{
        "headers": $utils.http.copyheaders($ctx.request.headers)
    }
}

""",
    response_mapping_template="""#if($ctx.result.statusCode == 200)
    $ctx.result.body
#else
    $utils.appendError($ctx.result.body, $ctx.result.statusCode)
#end

""")


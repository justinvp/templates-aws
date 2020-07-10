import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testGraphQLApi = new aws.appsync.GraphQLApi("test", {
    authenticationType: "API_KEY",
    schema: `type Mutation {
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
`,
});
const testDataSource = new aws.appsync.DataSource("test", {
    apiId: testGraphQLApi.id,
    httpConfig: {
        endpoint: "http://example.com",
    },
    type: "HTTP",
});
// UNIT type resolver (default)
const testResolver = new aws.appsync.Resolver("test", {
    apiId: testGraphQLApi.id,
    cachingConfig: {
        cachingKeys: [
            "$context.identity.sub",
            "$context.arguments.id",
        ],
        ttl: 60,
    },
    dataSource: testDataSource.name,
    field: "singlePost",
    requestTemplate: `{
    "version": "2018-05-29",
    "method": "GET",
    "resourcePath": "/",
    "params":{
        "headers": $utils.http.copyheaders($ctx.request.headers)
    }
}
`,
    responseTemplate: `#if($ctx.result.statusCode == 200)
    $ctx.result.body
#else
    $utils.appendError($ctx.result.body, $ctx.result.statusCode)
#end
`,
    type: "Query",
});
// PIPELINE type resolver
const mutationPipelineTest = new aws.appsync.Resolver("Mutation_pipelineTest", {
    apiId: testGraphQLApi.id,
    field: "pipelineTest",
    kind: "PIPELINE",
    pipelineConfig: {
        functions: [
            aws_appsync_function_test1.functionId,
            aws_appsync_function_test2.functionId,
            aws_appsync_function_test3.functionId,
        ],
    },
    requestTemplate: "{}",
    responseTemplate: "$util.toJson($ctx.result)",
    type: "Mutation",
});


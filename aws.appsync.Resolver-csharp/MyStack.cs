using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var testGraphQLApi = new Aws.AppSync.GraphQLApi("testGraphQLApi", new Aws.AppSync.GraphQLApiArgs
        {
            AuthenticationType = "API_KEY",
            Schema = @"type Mutation {
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

",
        });
        var testDataSource = new Aws.AppSync.DataSource("testDataSource", new Aws.AppSync.DataSourceArgs
        {
            ApiId = testGraphQLApi.Id,
            HttpConfig = new Aws.AppSync.Inputs.DataSourceHttpConfigArgs
            {
                Endpoint = "http://example.com",
            },
            Type = "HTTP",
        });
        // UNIT type resolver (default)
        var testResolver = new Aws.AppSync.Resolver("testResolver", new Aws.AppSync.ResolverArgs
        {
            ApiId = testGraphQLApi.Id,
            CachingConfig = new Aws.AppSync.Inputs.ResolverCachingConfigArgs
            {
                CachingKeys = 
                {
                    "$context.identity.sub",
                    "$context.arguments.id",
                },
                Ttl = 60,
            },
            DataSource = testDataSource.Name,
            Field = "singlePost",
            RequestTemplate = @"{
    ""version"": ""2018-05-29"",
    ""method"": ""GET"",
    ""resourcePath"": ""/"",
    ""params"":{
        ""headers"": $utils.http.copyheaders($ctx.request.headers)
    }
}

",
            ResponseTemplate = @"#if($ctx.result.statusCode == 200)
    $ctx.result.body
#else
    $utils.appendError($ctx.result.body, $ctx.result.statusCode)
#end

",
            Type = "Query",
        });
        // PIPELINE type resolver
        var mutationPipelineTest = new Aws.AppSync.Resolver("mutationPipelineTest", new Aws.AppSync.ResolverArgs
        {
            ApiId = testGraphQLApi.Id,
            Field = "pipelineTest",
            Kind = "PIPELINE",
            PipelineConfig = new Aws.AppSync.Inputs.ResolverPipelineConfigArgs
            {
                Functions = 
                {
                    aws_appsync_function.Test1.Function_id,
                    aws_appsync_function.Test2.Function_id,
                    aws_appsync_function.Test3.Function_id,
                },
            },
            RequestTemplate = "{}",
            ResponseTemplate = "$util.toJson($ctx.result)",
            Type = "Mutation",
        });
    }

}


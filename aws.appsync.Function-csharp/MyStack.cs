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
        var testFunction = new Aws.AppSync.Function("testFunction", new Aws.AppSync.FunctionArgs
        {
            ApiId = testGraphQLApi.Id,
            DataSource = testDataSource.Name,
            Name = "tf_example",
            RequestMappingTemplate = @"{
    ""version"": ""2018-05-29"",
    ""method"": ""GET"",
    ""resourcePath"": ""/"",
    ""params"":{
        ""headers"": $utils.http.copyheaders($ctx.request.headers)
    }
}

",
            ResponseMappingTemplate = @"#if($ctx.result.statusCode == 200)
    $ctx.result.body
#else
    $utils.appendError($ctx.result.body, $ctx.result.statusCode)
#end

",
        });
    }

}


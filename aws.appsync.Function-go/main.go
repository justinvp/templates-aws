package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/appsync"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testGraphQLApi, err := appsync.NewGraphQLApi(ctx, "testGraphQLApi", &appsync.GraphQLApiArgs{
			AuthenticationType: pulumi.String("API_KEY"),
			Schema:             pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "type Mutation {\n", "    putPost(id: ID!, title: String!): Post\n", "}\n", "\n", "type Post {\n", "    id: ID!\n", "    title: String!\n", "}\n", "\n", "type Query {\n", "    singlePost(id: ID!): Post\n", "}\n", "\n", "schema {\n", "    query: Query\n", "    mutation: Mutation\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		testDataSource, err := appsync.NewDataSource(ctx, "testDataSource", &appsync.DataSourceArgs{
			ApiId: testGraphQLApi.ID(),
			HttpConfig: &appsync.DataSourceHttpConfigArgs{
				Endpoint: pulumi.String("http://example.com"),
			},
			Type: pulumi.String("HTTP"),
		})
		if err != nil {
			return err
		}
		_, err = appsync.NewFunction(ctx, "testFunction", &appsync.FunctionArgs{
			ApiId:                   testGraphQLApi.ID(),
			DataSource:              testDataSource.Name,
			Name:                    pulumi.String("tf_example"),
			RequestMappingTemplate:  pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"version\": \"2018-05-29\",\n", "    \"method\": \"GET\",\n", "    \"resourcePath\": \"/\",\n", "    \"params\":{\n", "        \"headers\": ", "$", "utils.http.copyheaders(", "$", "ctx.request.headers)\n", "    }\n", "}\n", "\n")),
			ResponseMappingTemplate: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "#if(", "$", "ctx.result.statusCode == 200)\n", "    ", "$", "ctx.result.body\n", "#else\n", "    ", "$", "utils.appendError(", "$", "ctx.result.body, ", "$", "ctx.result.statusCode)\n", "#end\n", "\n")),
		})
		if err != nil {
			return err
		}
		return nil
	})
}


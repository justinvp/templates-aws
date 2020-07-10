package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleRestApi, err := apigateway.NewRestApi(ctx, "exampleRestApi", nil)
		if err != nil {
			return err
		}
		_, err = apigateway.NewDocumentationVersion(ctx, "exampleDocumentationVersion", &apigateway.DocumentationVersionArgs{
			Description: pulumi.String("Example description"),
			RestApiId:   exampleRestApi.ID(),
			Version:     pulumi.String("example_version"),
		}, pulumi.DependsOn([]pulumi.Resource{
			"aws_api_gateway_documentation_part.example",
		}))
		if err != nil {
			return err
		}
		_, err = apigateway.NewDocumentationPart(ctx, "exampleDocumentationPart", &apigateway.DocumentationPartArgs{
			Location: &apigateway.DocumentationPartLocationArgs{
				Type: pulumi.String("API"),
			},
			Properties: pulumi.String("{\"description\":\"Example\"}"),
			RestApiId:  exampleRestApi.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}


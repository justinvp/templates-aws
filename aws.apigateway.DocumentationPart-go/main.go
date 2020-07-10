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
		_, err = apigateway.NewDocumentationPart(ctx, "exampleDocumentationPart", &apigateway.DocumentationPartArgs{
			Location: &apigateway.DocumentationPartLocationArgs{
				Method: pulumi.String("GET"),
				Path:   pulumi.String("/example"),
				Type:   pulumi.String("METHOD"),
			},
			Properties: pulumi.String("{\"description\":\"Example description\"}"),
			RestApiId:  exampleRestApi.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}


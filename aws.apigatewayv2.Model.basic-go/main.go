package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigatewayv2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apigatewayv2.NewModel(ctx, "example", &apigatewayv2.ModelArgs{
			ApiId:       pulumi.Any(aws_apigatewayv2_api.Example.Id),
			ContentType: pulumi.String("application/json"),
			Schema:      pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"", "$", "schema\": \"http://json-schema.org/draft-04/schema#\",\n", "  \"title\": \"ExampleModel\",\n", "  \"type\": \"object\",\n", "  \"properties\": {\n", "    \"id\": { \"type\": \"string\" }\n", "  }\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		return nil
	})
}


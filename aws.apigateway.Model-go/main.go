package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myDemoAPI, err := apigateway.NewRestApi(ctx, "myDemoAPI", &apigateway.RestApiArgs{
			Description: pulumi.String("This is my API for demonstration purposes"),
		})
		if err != nil {
			return err
		}
		_, err = apigateway.NewModel(ctx, "myDemoModel", &apigateway.ModelArgs{
			ContentType: pulumi.String("application/json"),
			Description: pulumi.String("a JSON schema"),
			RestApi:     myDemoAPI.ID(),
			Schema:      pulumi.String(fmt.Sprintf("%v%v%v%v", "{\n", "  \"type\": \"object\"\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		return nil
	})
}


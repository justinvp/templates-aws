package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		main, err := apigateway.NewRestApi(ctx, "main", nil)
		if err != nil {
			return err
		}
		_, err = apigateway.NewResponse(ctx, "test", &apigateway.ResponseArgs{
			ResponseParameters: pulumi.StringMap{
				"gatewayresponse.header.Authorization": pulumi.String("'Basic'"),
			},
			ResponseTemplates: pulumi.StringMap{
				"application/json": pulumi.String(fmt.Sprintf("%v%v%v", "{'message':", "$", "context.error.messageString}")),
			},
			ResponseType: pulumi.String("UNAUTHORIZED"),
			RestApiId:    main.ID(),
			StatusCode:   pulumi.String("401"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}


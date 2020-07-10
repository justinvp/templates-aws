package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apigateway.NewRestApi(ctx, "myDemoAPI", &apigateway.RestApiArgs{
			Description: pulumi.String("This is my API for demonstration purposes"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}


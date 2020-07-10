package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myRestApi, err := apigateway.LookupRestApi(ctx, &apigateway.LookupRestApiArgs{
			Name: "my-rest-api",
		}, nil)
		if err != nil {
			return err
		}
		_, err = apigateway.LookupResource(ctx, &apigateway.LookupResourceArgs{
			Path:      "/endpoint/path",
			RestApiId: myRestApi.Id,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}


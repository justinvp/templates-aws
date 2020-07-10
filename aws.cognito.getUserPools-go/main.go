package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigateway"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cognito"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		selectedRestApi, err := apigateway.LookupRestApi(ctx, &apigateway.LookupRestApiArgs{
			Name: _var.Api_gateway_name,
		}, nil)
		if err != nil {
			return err
		}
		selectedUserPools, err := cognito.GetUserPools(ctx, &cognito.GetUserPoolsArgs{
			Name: _var.Cognito_user_pool_name,
		}, nil)
		if err != nil {
			return err
		}
		_, err = apigateway.NewAuthorizer(ctx, "cognito", &apigateway.AuthorizerArgs{
			ProviderArns: toPulumiStringArray(selectedUserPools.Arns),
			RestApi:      pulumi.String(selectedRestApi.Id),
			Type:         pulumi.String("COGNITO_USER_POOLS"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
func toPulumiStringArray(arr []string) pulumi.StringArray {
	var pulumiArr pulumi.StringArray
	for _, v := range arr {
		pulumiArr = append(pulumiArr, pulumi.String(v))
	}
	return pulumiArr
}


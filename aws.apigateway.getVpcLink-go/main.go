package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apigateway.LookupVpcLink(ctx, &apigateway.LookupVpcLinkArgs{
			Name: "my-vpc-link",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}


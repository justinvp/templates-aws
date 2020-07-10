package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigatewayv2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apigatewayv2.NewVpcLink(ctx, "example", &apigatewayv2.VpcLinkArgs{
			SecurityGroupIds: pulumi.StringArray{
				pulumi.Any(data.Aws_security_group.Example.Id),
			},
			SubnetIds: data.Aws_subnet_ids.Example.Ids,
			Tags: pulumi.StringMap{
				"Usage": pulumi.String("example"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}


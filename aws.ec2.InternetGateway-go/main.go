package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewInternetGateway(ctx, "gw", &ec2.InternetGatewayArgs{
			Tags: pulumi.StringMap{
				"Name": pulumi.String("main"),
			},
			VpcId: pulumi.Any(aws_vpc.Main.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}


package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/globalaccelerator"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := globalaccelerator.NewEndpointGroup(ctx, "example", &globalaccelerator.EndpointGroupArgs{
			EndpointConfigurations: globalaccelerator.EndpointGroupEndpointConfigurationArray{
				&globalaccelerator.EndpointGroupEndpointConfigurationArgs{
					EndpointId: pulumi.Any(aws_lb.Example.Arn),
					Weight:     pulumi.Int(100),
				},
			},
			ListenerArn: pulumi.Any(aws_globalaccelerator_listener.Example.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}


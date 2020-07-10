package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/appmesh"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appmesh.NewVirtualRouter(ctx, "serviceb", &appmesh.VirtualRouterArgs{
			MeshName: pulumi.Any(aws_appmesh_mesh.Simple.Id),
			Spec: &appmesh.VirtualRouterSpecArgs{
				Listener: &appmesh.VirtualRouterSpecListenerArgs{
					PortMapping: &appmesh.VirtualRouterSpecListenerPortMappingArgs{
						Port:     pulumi.Int(8080),
						Protocol: pulumi.String("http"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}


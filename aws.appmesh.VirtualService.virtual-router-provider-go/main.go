package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/appmesh"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appmesh.NewVirtualService(ctx, "servicea", &appmesh.VirtualServiceArgs{
			MeshName: pulumi.Any(aws_appmesh_mesh.Simple.Id),
			Spec: &appmesh.VirtualServiceSpecArgs{
				Provider: &appmesh.VirtualServiceSpecProviderArgs{
					VirtualRouter: &appmesh.VirtualServiceSpecProviderVirtualRouterArgs{
						VirtualRouterName: pulumi.Any(aws_appmesh_virtual_router.Serviceb.Name),
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


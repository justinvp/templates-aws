package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/appmesh"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appmesh.NewRoute(ctx, "serviceb", &appmesh.RouteArgs{
			MeshName: pulumi.Any(aws_appmesh_mesh.Simple.Id),
			Spec: &appmesh.RouteSpecArgs{
				TcpRoute: &appmesh.RouteSpecTcpRouteArgs{
					Action: &appmesh.RouteSpecTcpRouteActionArgs{
						WeightedTarget: pulumi.MapArray{
							pulumi.Map{
								"virtualNode": pulumi.Any(aws_appmesh_virtual_node.Serviceb1.Name),
								"weight":      pulumi.Float64(100),
							},
						},
					},
				},
			},
			VirtualRouterName: pulumi.Any(aws_appmesh_virtual_router.Serviceb.Name),
		})
		if err != nil {
			return err
		}
		return nil
	})
}


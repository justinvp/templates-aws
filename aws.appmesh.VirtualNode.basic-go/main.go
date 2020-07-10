package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/appmesh"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appmesh.NewVirtualNode(ctx, "serviceb1", &appmesh.VirtualNodeArgs{
			MeshName: pulumi.Any(aws_appmesh_mesh.Simple.Id),
			Spec: &appmesh.VirtualNodeSpecArgs{
				Backend: pulumi.StringMapMapArray{
					pulumi.StringMapMap{
						"virtualService": pulumi.StringMap{
							"virtualServiceName": pulumi.String("servicea.simpleapp.local"),
						},
					},
				},
				Listener: &appmesh.VirtualNodeSpecListenerArgs{
					PortMapping: &appmesh.VirtualNodeSpecListenerPortMappingArgs{
						Port:     pulumi.Int(8080),
						Protocol: pulumi.String("http"),
					},
				},
				ServiceDiscovery: &appmesh.VirtualNodeSpecServiceDiscoveryArgs{
					Dns: &appmesh.VirtualNodeSpecServiceDiscoveryDnsArgs{
						Hostname: pulumi.String("serviceb.simpleapp.local"),
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


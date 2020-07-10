package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/appmesh"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appmesh.NewMesh(ctx, "simple", &appmesh.MeshArgs{
			Spec: &appmesh.MeshSpecArgs{
				EgressFilter: &appmesh.MeshSpecEgressFilterArgs{
					Type: pulumi.String("ALLOW_ALL"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}


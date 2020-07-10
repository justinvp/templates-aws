package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/efs"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := efs.NewFileSystem(ctx, "fooWithLifecylePolicy", &efs.FileSystemArgs{
			LifecyclePolicy: &efs.FileSystemLifecyclePolicyArgs{
				TransitionToIa: pulumi.String("AFTER_30_DAYS"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}


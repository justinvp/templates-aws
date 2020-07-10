package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		production, err := ssm.NewPatchBaseline(ctx, "production", &ssm.PatchBaselineArgs{
			ApprovedPatches: pulumi.StringArray{
				pulumi.String("KB123456"),
			},
		})
		if err != nil {
			return err
		}
		_, err = ssm.NewPatchGroup(ctx, "patchgroup", &ssm.PatchGroupArgs{
			BaselineId: production.ID(),
			PatchGroup: pulumi.String("patch-group-name"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}


package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ebs"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ebs.GetSnapshotIds(ctx, &ebs.GetSnapshotIdsArgs{
			Filters: []ebs.GetSnapshotIdsFilter{
				ebs.GetSnapshotIdsFilter{
					Name: "volume-size",
					Values: []string{
						"40",
					},
				},
				ebs.GetSnapshotIdsFilter{
					Name: "tag:Name",
					Values: []string{
						"Example",
					},
				},
			},
			Owners: []string{
				"self",
			},
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}


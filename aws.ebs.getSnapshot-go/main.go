package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ebs"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := true
		_, err := ebs.LookupSnapshot(ctx, &ebs.LookupSnapshotArgs{
			Filters: []ebs.GetSnapshotFilter{
				ebs.GetSnapshotFilter{
					Name: "volume-size",
					Values: []string{
						"40",
					},
				},
				ebs.GetSnapshotFilter{
					Name: "tag:Name",
					Values: []string{
						"Example",
					},
				},
			},
			MostRecent: &opt0,
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


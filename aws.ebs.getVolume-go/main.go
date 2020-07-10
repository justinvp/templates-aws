package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ebs"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := true
		_, err := ebs.LookupVolume(ctx, &ebs.LookupVolumeArgs{
			Filters: []ebs.GetVolumeFilter{
				ebs.GetVolumeFilter{
					Name: "volume-type",
					Values: []string{
						"gp2",
					},
				},
				ebs.GetVolumeFilter{
					Name: "tag:Name",
					Values: []string{
						"Example",
					},
				},
			},
			MostRecent: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}


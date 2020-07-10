package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudtrail"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudtrail.NewTrail(ctx, "example", &cloudtrail.TrailArgs{
			EventSelectors: cloudtrail.TrailEventSelectorArray{
				&cloudtrail.TrailEventSelectorArgs{
					DataResource: pulumi.MapArray{
						pulumi.Map{
							"type": pulumi.String("AWS::Lambda::Function"),
							"values": pulumi.StringArray{
								pulumi.String("arn:aws:lambda"),
							},
						},
					},
					IncludeManagementEvents: pulumi.Bool(true),
					ReadWriteType:           pulumi.String("All"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}


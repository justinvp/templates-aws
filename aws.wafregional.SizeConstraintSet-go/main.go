package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/wafregional"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := wafregional.NewSizeConstraintSet(ctx, "sizeConstraintSet", &wafregional.SizeConstraintSetArgs{
			SizeConstraints: wafregional.SizeConstraintSetSizeConstraintArray{
				&wafregional.SizeConstraintSetSizeConstraintArgs{
					ComparisonOperator: pulumi.String("EQ"),
					FieldToMatch: &wafregional.SizeConstraintSetSizeConstraintFieldToMatchArgs{
						Type: pulumi.String("BODY"),
					},
					Size:               pulumi.Int(4096),
					TextTransformation: pulumi.String("NONE"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}


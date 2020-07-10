package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/waf"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := waf.NewSizeConstraintSet(ctx, "sizeConstraintSet", &waf.SizeConstraintSetArgs{
			SizeConstraints: waf.SizeConstraintSetSizeConstraintArray{
				&waf.SizeConstraintSetSizeConstraintArgs{
					ComparisonOperator: pulumi.String("EQ"),
					FieldToMatch: &waf.SizeConstraintSetSizeConstraintFieldToMatchArgs{
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


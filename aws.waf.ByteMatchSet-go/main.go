package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/waf"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := waf.NewByteMatchSet(ctx, "byteSet", &waf.ByteMatchSetArgs{
			ByteMatchTuples: waf.ByteMatchSetByteMatchTupleArray{
				&waf.ByteMatchSetByteMatchTupleArgs{
					FieldToMatch: &waf.ByteMatchSetByteMatchTupleFieldToMatchArgs{
						Data: pulumi.String("referer"),
						Type: pulumi.String("HEADER"),
					},
					PositionalConstraint: pulumi.String("CONTAINS"),
					TargetString:         pulumi.String("badrefer1"),
					TextTransformation:   pulumi.String("NONE"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}


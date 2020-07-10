package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/wafregional"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := wafregional.NewGeoMatchSet(ctx, "geoMatchSet", &wafregional.GeoMatchSetArgs{
			GeoMatchConstraints: wafregional.GeoMatchSetGeoMatchConstraintArray{
				&wafregional.GeoMatchSetGeoMatchConstraintArgs{
					Type:  pulumi.String("Country"),
					Value: pulumi.String("US"),
				},
				&wafregional.GeoMatchSetGeoMatchConstraintArgs{
					Type:  pulumi.String("Country"),
					Value: pulumi.String("CA"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}


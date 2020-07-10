package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/waf"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := waf.NewGeoMatchSet(ctx, "geoMatchSet", &waf.GeoMatchSetArgs{
			GeoMatchConstraints: waf.GeoMatchSetGeoMatchConstraintArray{
				&waf.GeoMatchSetGeoMatchConstraintArgs{
					Type:  pulumi.String("Country"),
					Value: pulumi.String("US"),
				},
				&waf.GeoMatchSetGeoMatchConstraintArgs{
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


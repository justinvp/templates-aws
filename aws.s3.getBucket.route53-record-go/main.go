package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/route53"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		selected, err := s3.LookupBucket(ctx, &s3.LookupBucketArgs{
			Bucket: "bucket.test.com",
		}, nil)
		if err != nil {
			return err
		}
		opt0 := "test.com."
		testZone, err := route53.LookupZone(ctx, &route53.LookupZoneArgs{
			Name: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		_, err = route53.NewRecord(ctx, "example", &route53.RecordArgs{
			Aliases: route53.RecordAliasArray{
				&route53.RecordAliasArgs{
					Name:   pulumi.String(selected.WebsiteDomain),
					ZoneId: pulumi.String(selected.HostedZoneId),
				},
			},
			Name:   pulumi.String("bucket"),
			Type:   pulumi.String("A"),
			ZoneId: pulumi.String(testZone.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}


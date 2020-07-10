package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudfront"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bucket, err := s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
			Acl: pulumi.String("private"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("My bucket"),
			},
		})
		if err != nil {
			return err
		}
		s3OriginId := "myS3Origin"
		_, err = cloudfront.NewDistribution(ctx, "s3Distribution", &cloudfront.DistributionArgs{
			Aliases: pulumi.StringArray{
				pulumi.String("mysite.example.com"),
				pulumi.String("yoursite.example.com"),
			},
			Comment: pulumi.String("Some comment"),
			DefaultCacheBehavior: &cloudfront.DistributionDefaultCacheBehaviorArgs{
				AllowedMethods: pulumi.StringArray{
					pulumi.String("DELETE"),
					pulumi.String("GET"),
					pulumi.String("HEAD"),
					pulumi.String("OPTIONS"),
					pulumi.String("PATCH"),
					pulumi.String("POST"),
					pulumi.String("PUT"),
				},
				CachedMethods: pulumi.StringArray{
					pulumi.String("GET"),
					pulumi.String("HEAD"),
				},
				DefaultTtl: pulumi.Int(3600),
				ForwardedValues: &cloudfront.DistributionDefaultCacheBehaviorForwardedValuesArgs{
					Cookies: &cloudfront.DistributionDefaultCacheBehaviorForwardedValuesCookiesArgs{
						Forward: pulumi.String("none"),
					},
					QueryString: pulumi.Bool(false),
				},
				MaxTtl:               pulumi.Int(86400),
				MinTtl:               pulumi.Int(0),
				TargetOriginId:       pulumi.String(s3OriginId),
				ViewerProtocolPolicy: pulumi.String("allow-all"),
			},
			DefaultRootObject: pulumi.String("index.html"),
			Enabled:           pulumi.Bool(true),
			IsIpv6Enabled:     pulumi.Bool(true),
			LoggingConfig: &cloudfront.DistributionLoggingConfigArgs{
				Bucket:         pulumi.String("mylogs.s3.amazonaws.com"),
				IncludeCookies: pulumi.Bool(false),
				Prefix:         pulumi.String("myprefix"),
			},
			OrderedCacheBehaviors: cloudfront.DistributionOrderedCacheBehaviorArray{
				&cloudfront.DistributionOrderedCacheBehaviorArgs{
					AllowedMethods: pulumi.StringArray{
						pulumi.String("GET"),
						pulumi.String("HEAD"),
						pulumi.String("OPTIONS"),
					},
					CachedMethods: pulumi.StringArray{
						pulumi.String("GET"),
						pulumi.String("HEAD"),
						pulumi.String("OPTIONS"),
					},
					Compress:   pulumi.Bool(true),
					DefaultTtl: pulumi.Int(86400),
					ForwardedValues: &cloudfront.DistributionOrderedCacheBehaviorForwardedValuesArgs{
						Cookies: &cloudfront.DistributionOrderedCacheBehaviorForwardedValuesCookiesArgs{
							Forward: pulumi.String("none"),
						},
						Headers: pulumi.StringArray{
							pulumi.String("Origin"),
						},
						QueryString: pulumi.Bool(false),
					},
					MaxTtl:               pulumi.Int(31536000),
					MinTtl:               pulumi.Int(0),
					PathPattern:          pulumi.String("/content/immutable/*"),
					TargetOriginId:       pulumi.String(s3OriginId),
					ViewerProtocolPolicy: pulumi.String("redirect-to-https"),
				},
				&cloudfront.DistributionOrderedCacheBehaviorArgs{
					AllowedMethods: pulumi.StringArray{
						pulumi.String("GET"),
						pulumi.String("HEAD"),
						pulumi.String("OPTIONS"),
					},
					CachedMethods: pulumi.StringArray{
						pulumi.String("GET"),
						pulumi.String("HEAD"),
					},
					Compress:   pulumi.Bool(true),
					DefaultTtl: pulumi.Int(3600),
					ForwardedValues: &cloudfront.DistributionOrderedCacheBehaviorForwardedValuesArgs{
						Cookies: &cloudfront.DistributionOrderedCacheBehaviorForwardedValuesCookiesArgs{
							Forward: pulumi.String("none"),
						},
						QueryString: pulumi.Bool(false),
					},
					MaxTtl:               pulumi.Int(86400),
					MinTtl:               pulumi.Int(0),
					PathPattern:          pulumi.String("/content/*"),
					TargetOriginId:       pulumi.String(s3OriginId),
					ViewerProtocolPolicy: pulumi.String("redirect-to-https"),
				},
			},
			Origins: cloudfront.DistributionOriginArray{
				&cloudfront.DistributionOriginArgs{
					DomainName: bucket.BucketRegionalDomainName,
					OriginId:   pulumi.String(s3OriginId),
					S3OriginConfig: &cloudfront.DistributionOriginS3OriginConfigArgs{
						OriginAccessIdentity: pulumi.String("origin-access-identity/cloudfront/ABCDEFG1234567"),
					},
				},
			},
			PriceClass: pulumi.String("PriceClass_200"),
			Restrictions: &cloudfront.DistributionRestrictionsArgs{
				GeoRestriction: &cloudfront.DistributionRestrictionsGeoRestrictionArgs{
					Locations: pulumi.StringArray{
						pulumi.String("US"),
						pulumi.String("CA"),
						pulumi.String("GB"),
						pulumi.String("DE"),
					},
					RestrictionType: pulumi.String("whitelist"),
				},
			},
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("production"),
			},
			ViewerCertificate: &cloudfront.DistributionViewerCertificateArgs{
				CloudfrontDefaultCertificate: pulumi.Bool(true),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}


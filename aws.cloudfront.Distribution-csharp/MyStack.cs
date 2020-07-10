using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {
            Acl = "private",
            Tags = 
            {
                { "Name", "My bucket" },
            },
        });
        var s3OriginId = "myS3Origin";
        var s3Distribution = new Aws.CloudFront.Distribution("s3Distribution", new Aws.CloudFront.DistributionArgs
        {
            Aliases = 
            {
                "mysite.example.com",
                "yoursite.example.com",
            },
            Comment = "Some comment",
            DefaultCacheBehavior = new Aws.CloudFront.Inputs.DistributionDefaultCacheBehaviorArgs
            {
                AllowedMethods = 
                {
                    "DELETE",
                    "GET",
                    "HEAD",
                    "OPTIONS",
                    "PATCH",
                    "POST",
                    "PUT",
                },
                CachedMethods = 
                {
                    "GET",
                    "HEAD",
                },
                DefaultTtl = 3600,
                ForwardedValues = new Aws.CloudFront.Inputs.DistributionDefaultCacheBehaviorForwardedValuesArgs
                {
                    Cookies = new Aws.CloudFront.Inputs.DistributionDefaultCacheBehaviorForwardedValuesCookiesArgs
                    {
                        Forward = "none",
                    },
                    QueryString = false,
                },
                MaxTtl = 86400,
                MinTtl = 0,
                TargetOriginId = s3OriginId,
                ViewerProtocolPolicy = "allow-all",
            },
            DefaultRootObject = "index.html",
            Enabled = true,
            IsIpv6Enabled = true,
            LoggingConfig = new Aws.CloudFront.Inputs.DistributionLoggingConfigArgs
            {
                Bucket = "mylogs.s3.amazonaws.com",
                IncludeCookies = false,
                Prefix = "myprefix",
            },
            OrderedCacheBehaviors = 
            {
                new Aws.CloudFront.Inputs.DistributionOrderedCacheBehaviorArgs
                {
                    AllowedMethods = 
                    {
                        "GET",
                        "HEAD",
                        "OPTIONS",
                    },
                    CachedMethods = 
                    {
                        "GET",
                        "HEAD",
                        "OPTIONS",
                    },
                    Compress = true,
                    DefaultTtl = 86400,
                    ForwardedValues = new Aws.CloudFront.Inputs.DistributionOrderedCacheBehaviorForwardedValuesArgs
                    {
                        Cookies = new Aws.CloudFront.Inputs.DistributionOrderedCacheBehaviorForwardedValuesCookiesArgs
                        {
                            Forward = "none",
                        },
                        Headers = 
                        {
                            "Origin",
                        },
                        QueryString = false,
                    },
                    MaxTtl = 31536000,
                    MinTtl = 0,
                    PathPattern = "/content/immutable/*",
                    TargetOriginId = s3OriginId,
                    ViewerProtocolPolicy = "redirect-to-https",
                },
                new Aws.CloudFront.Inputs.DistributionOrderedCacheBehaviorArgs
                {
                    AllowedMethods = 
                    {
                        "GET",
                        "HEAD",
                        "OPTIONS",
                    },
                    CachedMethods = 
                    {
                        "GET",
                        "HEAD",
                    },
                    Compress = true,
                    DefaultTtl = 3600,
                    ForwardedValues = new Aws.CloudFront.Inputs.DistributionOrderedCacheBehaviorForwardedValuesArgs
                    {
                        Cookies = new Aws.CloudFront.Inputs.DistributionOrderedCacheBehaviorForwardedValuesCookiesArgs
                        {
                            Forward = "none",
                        },
                        QueryString = false,
                    },
                    MaxTtl = 86400,
                    MinTtl = 0,
                    PathPattern = "/content/*",
                    TargetOriginId = s3OriginId,
                    ViewerProtocolPolicy = "redirect-to-https",
                },
            },
            Origins = 
            {
                new Aws.CloudFront.Inputs.DistributionOriginArgs
                {
                    DomainName = bucket.BucketRegionalDomainName,
                    OriginId = s3OriginId,
                    S3OriginConfig = new Aws.CloudFront.Inputs.DistributionOriginS3OriginConfigArgs
                    {
                        OriginAccessIdentity = "origin-access-identity/cloudfront/ABCDEFG1234567",
                    },
                },
            },
            PriceClass = "PriceClass_200",
            Restrictions = new Aws.CloudFront.Inputs.DistributionRestrictionsArgs
            {
                GeoRestriction = new Aws.CloudFront.Inputs.DistributionRestrictionsGeoRestrictionArgs
                {
                    Locations = 
                    {
                        "US",
                        "CA",
                        "GB",
                        "DE",
                    },
                    RestrictionType = "whitelist",
                },
            },
            Tags = 
            {
                { "Environment", "production" },
            },
            ViewerCertificate = new Aws.CloudFront.Inputs.DistributionViewerCertificateArgs
            {
                CloudfrontDefaultCertificate = true,
            },
        });
    }

}


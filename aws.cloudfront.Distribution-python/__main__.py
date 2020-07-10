import pulumi
import pulumi_aws as aws

bucket = aws.s3.Bucket("bucket",
    acl="private",
    tags={
        "Name": "My bucket",
    })
s3_origin_id = "myS3Origin"
s3_distribution = aws.cloudfront.Distribution("s3Distribution",
    aliases=[
        "mysite.example.com",
        "yoursite.example.com",
    ],
    comment="Some comment",
    default_cache_behavior={
        "allowedMethods": [
            "DELETE",
            "GET",
            "HEAD",
            "OPTIONS",
            "PATCH",
            "POST",
            "PUT",
        ],
        "cachedMethods": [
            "GET",
            "HEAD",
        ],
        "defaultTtl": 3600,
        "forwardedValues": {
            "cookies": {
                "forward": "none",
            },
            "queryString": False,
        },
        "maxTtl": 86400,
        "minTtl": 0,
        "targetOriginId": s3_origin_id,
        "viewerProtocolPolicy": "allow-all",
    },
    default_root_object="index.html",
    enabled=True,
    is_ipv6_enabled=True,
    logging_config={
        "bucket": "mylogs.s3.amazonaws.com",
        "includeCookies": False,
        "prefix": "myprefix",
    },
    ordered_cache_behaviors=[
        {
            "allowedMethods": [
                "GET",
                "HEAD",
                "OPTIONS",
            ],
            "cachedMethods": [
                "GET",
                "HEAD",
                "OPTIONS",
            ],
            "compress": True,
            "defaultTtl": 86400,
            "forwardedValues": {
                "cookies": {
                    "forward": "none",
                },
                "headers": ["Origin"],
                "queryString": False,
            },
            "maxTtl": 31536000,
            "minTtl": 0,
            "pathPattern": "/content/immutable/*",
            "targetOriginId": s3_origin_id,
            "viewerProtocolPolicy": "redirect-to-https",
        },
        {
            "allowedMethods": [
                "GET",
                "HEAD",
                "OPTIONS",
            ],
            "cachedMethods": [
                "GET",
                "HEAD",
            ],
            "compress": True,
            "defaultTtl": 3600,
            "forwardedValues": {
                "cookies": {
                    "forward": "none",
                },
                "queryString": False,
            },
            "maxTtl": 86400,
            "minTtl": 0,
            "pathPattern": "/content/*",
            "targetOriginId": s3_origin_id,
            "viewerProtocolPolicy": "redirect-to-https",
        },
    ],
    origins=[{
        "domain_name": bucket.bucket_regional_domain_name,
        "originId": s3_origin_id,
        "s3OriginConfig": {
            "originAccessIdentity": "origin-access-identity/cloudfront/ABCDEFG1234567",
        },
    }],
    price_class="PriceClass_200",
    restrictions={
        "geoRestriction": {
            "locations": [
                "US",
                "CA",
                "GB",
                "DE",
            ],
            "restrictionType": "whitelist",
        },
    },
    tags={
        "Environment": "production",
    },
    viewer_certificate={
        "cloudfrontDefaultCertificate": True,
    })


import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bucket = new aws.s3.Bucket("b", {
    acl: "private",
    tags: {
        Name: "My bucket",
    },
});
const s3OriginId = "myS3Origin";
const s3Distribution = new aws.cloudfront.Distribution("s3_distribution", {
    aliases: [
        "mysite.example.com",
        "yoursite.example.com",
    ],
    comment: "Some comment",
    defaultCacheBehavior: {
        allowedMethods: [
            "DELETE",
            "GET",
            "HEAD",
            "OPTIONS",
            "PATCH",
            "POST",
            "PUT",
        ],
        cachedMethods: [
            "GET",
            "HEAD",
        ],
        defaultTtl: 3600,
        forwardedValues: {
            cookies: {
                forward: "none",
            },
            queryString: false,
        },
        maxTtl: 86400,
        minTtl: 0,
        targetOriginId: s3OriginId,
        viewerProtocolPolicy: "allow-all",
    },
    defaultRootObject: "index.html",
    enabled: true,
    isIpv6Enabled: true,
    loggingConfig: {
        bucket: "mylogs.s3.amazonaws.com",
        includeCookies: false,
        prefix: "myprefix",
    },
    orderedCacheBehaviors: [
        // Cache behavior with precedence 0
        {
            allowedMethods: [
                "GET",
                "HEAD",
                "OPTIONS",
            ],
            cachedMethods: [
                "GET",
                "HEAD",
                "OPTIONS",
            ],
            compress: true,
            defaultTtl: 86400,
            forwardedValues: {
                cookies: {
                    forward: "none",
                },
                headers: ["Origin"],
                queryString: false,
            },
            maxTtl: 31536000,
            minTtl: 0,
            pathPattern: "/content/immutable/*",
            targetOriginId: s3OriginId,
            viewerProtocolPolicy: "redirect-to-https",
        },
        // Cache behavior with precedence 1
        {
            allowedMethods: [
                "GET",
                "HEAD",
                "OPTIONS",
            ],
            cachedMethods: [
                "GET",
                "HEAD",
            ],
            compress: true,
            defaultTtl: 3600,
            forwardedValues: {
                cookies: {
                    forward: "none",
                },
                queryString: false,
            },
            maxTtl: 86400,
            minTtl: 0,
            pathPattern: "/content/*",
            targetOriginId: s3OriginId,
            viewerProtocolPolicy: "redirect-to-https",
        },
    ],
    origins: [{
        domainName: bucket.bucketRegionalDomainName,
        originId: s3OriginId,
        s3OriginConfig: {
            originAccessIdentity: "origin-access-identity/cloudfront/ABCDEFG1234567",
        },
    }],
    priceClass: "PriceClass_200",
    restrictions: {
        geoRestriction: {
            locations: [
                "US",
                "CA",
                "GB",
                "DE",
            ],
            restrictionType: "whitelist",
        },
    },
    tags: {
        Environment: "production",
    },
    viewerCertificate: {
        cloudfrontDefaultCertificate: true,
    },
});


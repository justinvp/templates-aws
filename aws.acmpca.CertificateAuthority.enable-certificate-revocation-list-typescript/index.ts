import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucket = new aws.s3.Bucket("example", {});
const acmpcaBucketAccess = pulumi.all([exampleBucket.arn, exampleBucket.arn]).apply(([exampleBucketArn, exampleBucketArn1]) => aws.iam.getPolicyDocument({
    statements: [{
        actions: [
            "s3:GetBucketAcl",
            "s3:GetBucketLocation",
            "s3:PutObject",
            "s3:PutObjectAcl",
        ],
        principals: [{
            identifiers: ["acm-pca.amazonaws.com"],
            type: "Service",
        }],
        resources: [
            exampleBucketArn,
            `${exampleBucketArn1}/*`,
        ],
    }],
}, { async: true }));
const exampleBucketPolicy = new aws.s3.BucketPolicy("example", {
    bucket: exampleBucket.id,
    policy: acmpcaBucketAccess.json,
});
const exampleCertificateAuthority = new aws.acmpca.CertificateAuthority("example", {
    certificateAuthorityConfiguration: {
        keyAlgorithm: "RSA_4096",
        signingAlgorithm: "SHA512WITHRSA",
        subject: {
            commonName: "example.com",
        },
    },
    revocationConfiguration: {
        crlConfiguration: {
            customCname: "crl.example.com",
            enabled: true,
            expirationInDays: 7,
            s3BucketName: exampleBucket.id,
        },
    },
}, { dependsOn: [exampleBucketPolicy] });


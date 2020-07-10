import pulumi
import pulumi_aws as aws

example_bucket = aws.s3.Bucket("exampleBucket")
acmpca_bucket_access = pulumi.Output.all(example_bucket.arn, example_bucket.arn).apply(lambda exampleBucketArn, exampleBucketArn1: aws.iam.get_policy_document(statements=[{
    "actions": [
        "s3:GetBucketAcl",
        "s3:GetBucketLocation",
        "s3:PutObject",
        "s3:PutObjectAcl",
    ],
    "principals": [{
        "identifiers": ["acm-pca.amazonaws.com"],
        "type": "Service",
    }],
    "resources": [
        example_bucket_arn,
        f"{example_bucket_arn1}/*",
    ],
}]))
example_bucket_policy = aws.s3.BucketPolicy("exampleBucketPolicy",
    bucket=example_bucket.id,
    policy=acmpca_bucket_access.json)
example_certificate_authority = aws.acmpca.CertificateAuthority("exampleCertificateAuthority",
    certificate_authority_configuration={
        "keyAlgorithm": "RSA_4096",
        "signingAlgorithm": "SHA512WITHRSA",
        "subject": {
            "commonName": "example.com",
        },
    },
    revocation_configuration={
        "crlConfiguration": {
            "customCname": "crl.example.com",
            "enabled": True,
            "expirationInDays": 7,
            "s3_bucket_name": example_bucket.id,
        },
    },
    opts=ResourceOptions(depends_on=["aws_s3_bucket_policy.example"]))


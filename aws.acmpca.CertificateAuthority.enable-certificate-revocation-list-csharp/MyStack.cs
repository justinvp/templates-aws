using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleBucket = new Aws.S3.Bucket("exampleBucket", new Aws.S3.BucketArgs
        {
        });
        var acmpcaBucketAccess = Output.Tuple(exampleBucket.Arn, exampleBucket.Arn).Apply(values =>
        {
            var exampleBucketArn = values.Item1;
            var exampleBucketArn1 = values.Item2;
            return Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
            {
                Statements = 
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementArgs
                    {
                        Actions = 
                        {
                            "s3:GetBucketAcl",
                            "s3:GetBucketLocation",
                            "s3:PutObject",
                            "s3:PutObjectAcl",
                        },
                        Principals = 
                        {
                            new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalArgs
                            {
                                Identifiers = 
                                {
                                    "acm-pca.amazonaws.com",
                                },
                                Type = "Service",
                            },
                        },
                        Resources = 
                        {
                            exampleBucketArn,
                            $"{exampleBucketArn1}/*",
                        },
                    },
                },
            });
        });
        var exampleBucketPolicy = new Aws.S3.BucketPolicy("exampleBucketPolicy", new Aws.S3.BucketPolicyArgs
        {
            Bucket = exampleBucket.Id,
            Policy = acmpcaBucketAccess.Apply(acmpcaBucketAccess => acmpcaBucketAccess.Json),
        });
        var exampleCertificateAuthority = new Aws.Acmpca.CertificateAuthority("exampleCertificateAuthority", new Aws.Acmpca.CertificateAuthorityArgs
        {
            CertificateAuthorityConfiguration = new Aws.Acmpca.Inputs.CertificateAuthorityCertificateAuthorityConfigurationArgs
            {
                KeyAlgorithm = "RSA_4096",
                SigningAlgorithm = "SHA512WITHRSA",
                Subject = new Aws.Acmpca.Inputs.CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs
                {
                    CommonName = "example.com",
                },
            },
            RevocationConfiguration = new Aws.Acmpca.Inputs.CertificateAuthorityRevocationConfigurationArgs
            {
                CrlConfiguration = new Aws.Acmpca.Inputs.CertificateAuthorityRevocationConfigurationCrlConfigurationArgs
                {
                    CustomCname = "crl.example.com",
                    Enabled = true,
                    ExpirationInDays = 7,
                    S3BucketName = exampleBucket.Id,
                },
            },
        }, new CustomResourceOptions
        {
            DependsOn = 
            {
                "aws_s3_bucket_policy.example",
            },
        });
    }

}


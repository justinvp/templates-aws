import pulumi
import pulumi_aws as aws

example = aws.acmpca.CertificateAuthority("example",
    certificate_authority_configuration={
        "keyAlgorithm": "RSA_4096",
        "signingAlgorithm": "SHA512WITHRSA",
        "subject": {
            "commonName": "example.com",
        },
    },
    permanent_deletion_time_in_days=7)


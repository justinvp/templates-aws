import pulumi
import pulumi_aws as aws

example = aws.acmpca.get_certificate_authority(arn="arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012")


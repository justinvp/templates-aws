import pulumi
import pulumi_aws as aws

cert_certificate = aws.acm.Certificate("certCertificate",
    domain_name="example.com",
    validation_method="EMAIL")
cert_certificate_validation = aws.acm.CertificateValidation("certCertificateValidation", certificate_arn=cert_certificate.arn)


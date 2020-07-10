import pulumi
import pulumi_aws as aws

cert_certificate = aws.acm.Certificate("certCertificate",
    domain_name="example.com",
    validation_method="DNS")
zone = aws.route53.get_zone(name="example.com.",
    private_zone=False)
cert_validation = aws.route53.Record("certValidation",
    name=cert_certificate.domain_validation_options[0]["resourceRecordName"],
    records=[cert_certificate.domain_validation_options[0]["resourceRecordValue"]],
    ttl=60,
    type=cert_certificate.domain_validation_options[0]["resourceRecordType"],
    zone_id=zone.zone_id)
cert_certificate_validation = aws.acm.CertificateValidation("certCertificateValidation",
    certificate_arn=cert_certificate.arn,
    validation_record_fqdns=[cert_validation.fqdn])
front_end = aws.lb.Listener("frontEnd", certificate_arn=cert_certificate_validation.certificate_arn)


import pulumi
import pulumi_aws as aws

cert_certificate = aws.acm.Certificate("certCertificate",
    domain_name="example.com",
    subject_alternative_names=[
        "www.example.com",
        "example.org",
    ],
    validation_method="DNS")
zone = aws.route53.get_zone(name="example.com.",
    private_zone=False)
zone_alt = aws.route53.get_zone(name="example.org.",
    private_zone=False)
cert_validation = aws.route53.Record("certValidation",
    name=cert_certificate.domain_validation_options[0]["resourceRecordName"],
    records=[cert_certificate.domain_validation_options[0]["resourceRecordValue"]],
    ttl=60,
    type=cert_certificate.domain_validation_options[0]["resourceRecordType"],
    zone_id=zone.zone_id)
cert_validation_alt1 = aws.route53.Record("certValidationAlt1",
    name=cert_certificate.domain_validation_options[1]["resourceRecordName"],
    records=[cert_certificate.domain_validation_options[1]["resourceRecordValue"]],
    ttl=60,
    type=cert_certificate.domain_validation_options[1]["resourceRecordType"],
    zone_id=zone.zone_id)
cert_validation_alt2 = aws.route53.Record("certValidationAlt2",
    name=cert_certificate.domain_validation_options[2]["resourceRecordName"],
    records=[cert_certificate.domain_validation_options[2]["resourceRecordValue"]],
    ttl=60,
    type=cert_certificate.domain_validation_options[2]["resourceRecordType"],
    zone_id=zone_alt.zone_id)
cert_certificate_validation = aws.acm.CertificateValidation("certCertificateValidation",
    certificate_arn=cert_certificate.arn,
    validation_record_fqdns=[
        cert_validation.fqdn,
        cert_validation_alt1.fqdn,
        cert_validation_alt2.fqdn,
    ])
front_end = aws.lb.Listener("frontEnd", certificate_arn=cert_certificate_validation.certificate_arn)


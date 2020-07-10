import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const certCertificate = new aws.acm.Certificate("cert", {
    domainName: "example.com",
    subjectAlternativeNames: [
        "www.example.com",
        "example.org",
    ],
    validationMethod: "DNS",
});
const zone = pulumi.output(aws.route53.getZone({
    name: "example.com.",
    privateZone: false,
}, { async: true }));
const zoneAlt = pulumi.output(aws.route53.getZone({
    name: "example.org.",
    privateZone: false,
}, { async: true }));
const certValidation = new aws.route53.Record("cert_validation", {
    name: certCertificate.domainValidationOptions[0].resourceRecordName,
    records: [certCertificate.domainValidationOptions[0].resourceRecordValue],
    ttl: 60,
    type: certCertificate.domainValidationOptions[0].resourceRecordType,
    zoneId: zone.zoneId!,
});
const certValidationAlt1 = new aws.route53.Record("cert_validation_alt1", {
    name: certCertificate.domainValidationOptions[1].resourceRecordName,
    records: [certCertificate.domainValidationOptions[1].resourceRecordValue],
    ttl: 60,
    type: certCertificate.domainValidationOptions[1].resourceRecordType,
    zoneId: zone.zoneId!,
});
const certValidationAlt2 = new aws.route53.Record("cert_validation_alt2", {
    name: certCertificate.domainValidationOptions[2].resourceRecordName,
    records: [certCertificate.domainValidationOptions[2].resourceRecordValue],
    ttl: 60,
    type: certCertificate.domainValidationOptions[2].resourceRecordType,
    zoneId: zoneAlt.zoneId!,
});
const certCertificateValidation = new aws.acm.CertificateValidation("cert", {
    certificateArn: certCertificate.arn,
    validationRecordFqdns: [
        certValidation.fqdn,
        certValidationAlt1.fqdn,
        certValidationAlt2.fqdn,
    ],
});
const frontEnd = new aws.lb.Listener("front_end", {
    // [...]
    certificateArn: certCertificateValidation.certificateArn,
});


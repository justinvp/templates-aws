import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const certCertificate = new aws.acm.Certificate("cert", {
    domainName: "example.com",
    validationMethod: "DNS",
});
const zone = pulumi.output(aws.route53.getZone({
    name: "example.com.",
    privateZone: false,
}, { async: true }));
const certValidation = new aws.route53.Record("cert_validation", {
    name: certCertificate.domainValidationOptions[0].resourceRecordName,
    records: [certCertificate.domainValidationOptions[0].resourceRecordValue],
    ttl: 60,
    type: certCertificate.domainValidationOptions[0].resourceRecordType,
    zoneId: zone.zoneId!,
});
const certCertificateValidation = new aws.acm.CertificateValidation("cert", {
    certificateArn: certCertificate.arn,
    validationRecordFqdns: [certValidation.fqdn],
});
const frontEnd = new aws.lb.Listener("front_end", {
    // [...]
    certificateArn: certCertificateValidation.certificateArn,
});


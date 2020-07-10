import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const certCertificate = new aws.acm.Certificate("cert", {
    domainName: "example.com",
    validationMethod: "EMAIL",
});
const certCertificateValidation = new aws.acm.CertificateValidation("cert", {
    certificateArn: certCertificate.arn,
});


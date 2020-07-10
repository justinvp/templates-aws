import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.acmpca.CertificateAuthority("example", {
    certificateAuthorityConfiguration: {
        keyAlgorithm: "RSA_4096",
        signingAlgorithm: "SHA512WITHRSA",
        subject: {
            commonName: "example.com",
        },
    },
    permanentDeletionTimeInDays: 7,
});


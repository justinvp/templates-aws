import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as tls from "@pulumi/tls";

const examplePrivateKey = new tls.PrivateKey("example", {
    algorithm: "RSA",
});
const exampleSelfSignedCert = new tls.SelfSignedCert("example", {
    allowedUses: [
        "key_encipherment",
        "digital_signature",
        "server_auth",
    ],
    keyAlgorithm: "RSA",
    privateKeyPem: examplePrivateKey.privateKeyPem,
    subjects: [{
        commonName: "example.com",
        organization: "ACME Examples, Inc",
    }],
    validityPeriodHours: 12,
});
const cert = new aws.acm.Certificate("cert", {
    certificateBody: exampleSelfSignedCert.certPem,
    privateKey: examplePrivateKey.privateKeyPem,
});


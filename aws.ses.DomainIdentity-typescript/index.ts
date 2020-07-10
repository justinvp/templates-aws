import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ses.DomainIdentity("example", {
    domain: "example.com",
});
const exampleAmazonsesVerificationRecord = new aws.route53.Record("example_amazonses_verification_record", {
    name: "_amazonses.example.com",
    records: [example.verificationToken],
    ttl: 600,
    type: "TXT",
    zoneId: "ABCDEFGHIJ123",
});


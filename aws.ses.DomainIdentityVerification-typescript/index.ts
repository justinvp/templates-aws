import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ses.DomainIdentity("example", {
    domain: "example.com",
});
const exampleAmazonsesVerificationRecord = new aws.route53.Record("example_amazonses_verification_record", {
    name: pulumi.interpolate`_amazonses.${example.id}`,
    records: [example.verificationToken],
    ttl: 600,
    type: "TXT",
    zoneId: aws_route53_zone_example.zoneId,
});
const exampleVerification = new aws.ses.DomainIdentityVerification("example_verification", {
    domain: example.id,
}, { dependsOn: [exampleAmazonsesVerificationRecord] });


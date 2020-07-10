import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleDomainName = new aws.apigateway.DomainName("example", {
    domainName: "api.example.com",
    endpointConfiguration: {
        types: "REGIONAL",
    },
    regionalCertificateArn: aws_acm_certificate_validation_example.certificateArn,
});
// Example DNS record using Route53.
// Route53 is not specifically required; any DNS host can be used.
const exampleRecord = new aws.route53.Record("example", {
    aliases: [{
        evaluateTargetHealth: true,
        name: exampleDomainName.regionalDomainName,
        zoneId: exampleDomainName.regionalZoneId,
    }],
    name: exampleDomainName.domainName,
    type: "A",
    zoneId: aws_route53_zone_example.id,
});


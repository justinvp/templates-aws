import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const selected = pulumi.output(aws.s3.getBucket({
    bucket: "bucket.test.com",
}, { async: true }));
const testZone = pulumi.output(aws.route53.getZone({
    name: "test.com.",
}, { async: true }));
const example = new aws.route53.Record("example", {
    aliases: [{
        name: selected.websiteDomain,
        zoneId: selected.hostedZoneId,
    }],
    name: "bucket",
    type: "A",
    zoneId: testZone.id,
});


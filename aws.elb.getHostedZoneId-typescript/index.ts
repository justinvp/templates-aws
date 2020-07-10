import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = pulumi.output(aws.elb.getHostedZoneId({ async: true }));
const www = new aws.route53.Record("www", {
    aliases: [{
        evaluateTargetHealth: true,
        name: aws_elb_main.dnsName,
        zoneId: main.id,
    }],
    name: "example.com",
    type: "A",
    zoneId: aws_route53_zone_primary.zoneId,
});


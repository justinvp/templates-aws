import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const www_dev = new aws.route53.Record("www-dev", {
    name: "www",
    records: ["dev.example.com"],
    setIdentifier: "dev",
    ttl: 5,
    type: "CNAME",
    weightedRoutingPolicies: [{
        weight: 10,
    }],
    zoneId: aws_route53_zone_primary.zoneId,
});
const www_live = new aws.route53.Record("www-live", {
    name: "www",
    records: ["live.example.com"],
    setIdentifier: "live",
    ttl: 5,
    type: "CNAME",
    weightedRoutingPolicies: [{
        weight: 90,
    }],
    zoneId: aws_route53_zone_primary.zoneId,
});


import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.elb.LoadBalancer("main", {
    availabilityZones: ["us-east-1c"],
    listeners: [{
        instancePort: 80,
        instanceProtocol: "http",
        lbPort: 80,
        lbProtocol: "http",
    }],
});
const www = new aws.route53.Record("www", {
    aliases: [{
        evaluateTargetHealth: true,
        name: main.dnsName,
        zoneId: main.zoneId,
    }],
    name: "example.com",
    type: "A",
    zoneId: aws_route53_zone_primary.zoneId,
});


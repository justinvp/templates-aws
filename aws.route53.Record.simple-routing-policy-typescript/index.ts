import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const www = new aws.route53.Record("www", {
    name: "www.example.com",
    records: [aws_eip_lb.publicIp],
    ttl: 300,
    type: "A",
    zoneId: aws_route53_zone_primary.zoneId,
});


import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.route53.Zone("main", {});
const dev = new aws.route53.Zone("dev", {
    tags: {
        Environment: "dev",
    },
});
const dev_ns = new aws.route53.Record("dev-ns", {
    name: "dev.example.com",
    records: [
        dev.nameServers[0],
        dev.nameServers[1],
        dev.nameServers[2],
        dev.nameServers[3],
    ],
    ttl: 30,
    type: "NS",
    zoneId: main.zoneId,
});


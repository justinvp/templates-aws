import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleZone = new aws.route53.Zone("example", {});
const exampleRecord = new aws.route53.Record("example", {
    allowOverwrite: true,
    name: "test.example.com",
    records: [
        exampleZone.nameServers[0],
        exampleZone.nameServers[1],
        exampleZone.nameServers[2],
        exampleZone.nameServers[3],
    ],
    ttl: 30,
    type: "NS",
    zoneId: exampleZone.zoneId,
});


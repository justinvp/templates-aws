import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const selected = pulumi.output(aws.route53.getZone({
    name: "test.com.",
    privateZone: true,
}, { async: true }));
const www = new aws.route53.Record("www", {
    name: pulumi.interpolate`www.${selected.name!}`,
    records: ["10.0.0.1"],
    ttl: 300,
    type: "A",
    zoneId: selected.zoneId!,
});


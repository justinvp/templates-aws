import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.route53.DelegationSet("main", {
    referenceName: "DynDNS",
});
const primary = new aws.route53.Zone("primary", {
    delegationSetId: main.id,
});
const secondary = new aws.route53.Zone("secondary", {
    delegationSetId: main.id,
});


import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.securityhub.Account("example", {});
const cis = new aws.securityhub.StandardsSubscription("cis", {
    standardsArn: "arn:aws:securityhub:::ruleset/cis-aws-foundations-benchmark/v/1.2.0",
}, { dependsOn: [example] });
const pci321 = new aws.securityhub.StandardsSubscription("pci_321", {
    standardsArn: "arn:aws:securityhub:us-east-1::standards/pci-dss/v/3.2.1",
}, { dependsOn: [example] });


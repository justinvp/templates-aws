import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.elasticsearch.Domain("example", {
    clusterConfig: {
        instanceType: "r4.large.elasticsearch",
    },
    elasticsearchVersion: "1.5",
    snapshotOptions: {
        automatedSnapshotStartHour: 23,
    },
    tags: {
        Domain: "TestDomain",
    },
});


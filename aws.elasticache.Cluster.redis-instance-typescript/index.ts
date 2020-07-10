import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.elasticache.Cluster("example", {
    engine: "redis",
    engineVersion: "3.2.10",
    nodeType: "cache.m4.large",
    numCacheNodes: 1,
    parameterGroupName: "default.redis3.2",
    port: 6379,
});


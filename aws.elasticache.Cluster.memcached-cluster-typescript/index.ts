import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.elasticache.Cluster("example", {
    engine: "memcached",
    nodeType: "cache.m4.large",
    numCacheNodes: 2,
    parameterGroupName: "default.memcached1.4",
    port: 11211,
});


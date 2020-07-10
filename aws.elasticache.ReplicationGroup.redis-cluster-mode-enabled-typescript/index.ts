import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const baz = new aws.elasticache.ReplicationGroup("baz", {
    automaticFailoverEnabled: true,
    clusterMode: {
        numNodeGroups: 2,
        replicasPerNodeGroup: 1,
    },
    nodeType: "cache.t2.small",
    parameterGroupName: "default.redis3.2.cluster.on",
    port: 6379,
    replicationGroupDescription: "test description",
});


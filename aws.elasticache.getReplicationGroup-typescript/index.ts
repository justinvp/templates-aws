import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bar = pulumi.output(aws.elasticache.getReplicationGroup({
    replicationGroupId: "example",
}, { async: true }));


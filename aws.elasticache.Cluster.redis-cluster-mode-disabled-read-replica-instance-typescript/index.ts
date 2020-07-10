import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const replica = new aws.elasticache.Cluster("replica", {
    replicationGroupId: aws_elasticache_replication_group_example.id,
});


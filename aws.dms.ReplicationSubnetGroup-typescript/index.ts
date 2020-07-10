import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new replication subnet group
const test = new aws.dms.ReplicationSubnetGroup("test", {
    replicationSubnetGroupDescription: "Test replication subnet group",
    replicationSubnetGroupId: "test-dms-replication-subnet-group-tf",
    subnetIds: ["subnet-12345678"],
    tags: {
        Name: "test",
    },
});


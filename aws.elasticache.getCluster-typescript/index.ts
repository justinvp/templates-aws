import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myCluster = pulumi.output(aws.elasticache.getCluster({
    clusterId: "my-cluster-id",
}, { async: true }));


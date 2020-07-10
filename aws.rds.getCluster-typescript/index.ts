import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const clusterName = pulumi.output(aws.rds.getCluster({
    clusterIdentifier: "clusterName",
}, { async: true }));


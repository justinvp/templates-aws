import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const cluster = pulumi.output(aws.cloudhsmv2.getCluster({
    clusterId: "cluster-testclusterid",
}, { async: true }));


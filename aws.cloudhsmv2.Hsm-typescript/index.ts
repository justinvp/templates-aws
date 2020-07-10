import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const cluster = pulumi.output(aws.cloudhsmv2.getCluster({
    clusterId: var_cloudhsm_cluster_id,
}, { async: true }));
const cloudhsmV2Hsm = new aws.cloudhsmv2.Hsm("cloudhsm_v2_hsm", {
    clusterId: cluster.clusterId,
    subnetId: cluster.apply(cluster => cluster.subnetIds[0]),
});


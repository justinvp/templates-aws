import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const task = new aws.emr.InstanceGroup("task", {
    clusterId: aws_emr_cluster_tf_test_cluster.id,
    instanceCount: 1,
    instanceType: "m5.xlarge",
});


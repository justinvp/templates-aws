import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bar = new aws.dax.Cluster("bar", {
    clusterName: "cluster-example",
    iamRoleArn: aws_iam_role_example.arn,
    nodeType: "dax.r4.large",
    replicationFactor: 1,
});


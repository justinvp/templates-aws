import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultCluster = new aws.redshift.Cluster("default", {
    clusterIdentifier: "tf-redshift-cluster",
    clusterType: "single-node",
    databaseName: "mydb",
    masterPassword: "Mustbe8characters",
    masterUsername: "foo",
    nodeType: "dc1.large",
});


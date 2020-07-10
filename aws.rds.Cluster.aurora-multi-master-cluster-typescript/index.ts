import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.rds.Cluster("example", {
    clusterIdentifier: "example",
    dbSubnetGroupName: aws_db_subnet_group_example.name,
    engineMode: "multimaster",
    masterPassword: "barbarbarbar",
    masterUsername: "foo",
    skipFinalSnapshot: true,
});


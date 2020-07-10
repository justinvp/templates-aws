import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bar = new aws.rds.Instance("bar", {
    allocatedStorage: 10,
    backupRetentionPeriod: 0,
    engine: "MySQL",
    engineVersion: "5.6.21",
    instanceClass: "db.t2.micro",
    maintenanceWindow: "Fri:09:00-Fri:09:30",
    name: "baz",
    parameterGroupName: "default.mysql5.6",
    password: "barbarbarbar",
    username: "foo",
});
const test = new aws.rds.Snapshot("test", {
    dbInstanceIdentifier: bar.id,
    dbSnapshotIdentifier: "testsnapshot1234",
});


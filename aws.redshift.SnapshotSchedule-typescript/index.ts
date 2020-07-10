import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultSnapshotSchedule = new aws.redshift.SnapshotSchedule("default", {
    definitions: ["rate(12 hours)"],
    identifier: "tf-redshift-snapshot-schedule",
});


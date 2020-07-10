import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.backup.getPlan({
    planId: "tf_example_backup_plan_id",
}, { async: true }));


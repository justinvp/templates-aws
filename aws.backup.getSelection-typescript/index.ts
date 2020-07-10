import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws_backup_plan_example.id.apply(id => aws.backup.getSelection({
    planId: id,
    selectionId: "selection-id-example",
}, { async: true }));


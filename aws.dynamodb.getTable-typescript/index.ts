import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const tableName = pulumi.output(aws.dynamodb.getTable({
    name: "tableName",
}, { async: true }));


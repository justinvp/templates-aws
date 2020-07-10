import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const database = pulumi.output(aws.rds.getInstance({
    dbInstanceIdentifier: "my-test-database",
}, { async: true }));


import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.qldb.getLedger({
    name: "an_example_ledger",
}, { async: true }));


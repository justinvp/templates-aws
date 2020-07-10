import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const reportDefinition = pulumi.output(aws.cur.getReportDefinition({
    reportName: "example",
}, { async: true }));


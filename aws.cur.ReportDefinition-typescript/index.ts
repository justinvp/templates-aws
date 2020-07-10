import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleCurReportDefinition = new aws.cur.ReportDefinition("example_cur_report_definition", {
    additionalArtifacts: [
        "REDSHIFT",
        "QUICKSIGHT",
    ],
    additionalSchemaElements: ["RESOURCES"],
    compression: "GZIP",
    format: "textORcsv",
    reportName: "example-cur-report-definition",
    s3Bucket: "example-bucket-name",
    s3Region: "us-east-1",
    timeUnit: "HOURLY",
});


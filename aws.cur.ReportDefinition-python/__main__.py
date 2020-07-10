import pulumi
import pulumi_aws as aws

example_cur_report_definition = aws.cur.ReportDefinition("exampleCurReportDefinition",
    additional_artifacts=[
        "REDSHIFT",
        "QUICKSIGHT",
    ],
    additional_schema_elements=["RESOURCES"],
    compression="GZIP",
    format="textORcsv",
    report_name="example-cur-report-definition",
    s3_bucket="example-bucket-name",
    s3_region="us-east-1",
    time_unit="HOURLY")


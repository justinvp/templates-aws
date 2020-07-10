import pulumi
import pulumi_aws as aws

report_definition = aws.cur.get_report_definition(report_name="example")


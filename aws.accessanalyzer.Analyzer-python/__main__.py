import pulumi
import pulumi_aws as aws

example = aws.accessanalyzer.Analyzer("example", analyzer_name="example")


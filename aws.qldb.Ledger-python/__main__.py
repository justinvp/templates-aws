import pulumi
import pulumi_aws as aws

sample_ledger = aws.qldb.Ledger("sample-ledger")


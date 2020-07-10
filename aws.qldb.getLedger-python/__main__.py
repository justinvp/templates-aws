import pulumi
import pulumi_aws as aws

example = aws.qldb.get_ledger(name="an_example_ledger")


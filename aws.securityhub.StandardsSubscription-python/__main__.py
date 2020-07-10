import pulumi
import pulumi_aws as aws

example = aws.securityhub.Account("example")
cis = aws.securityhub.StandardsSubscription("cis", standards_arn="arn:aws:securityhub:::ruleset/cis-aws-foundations-benchmark/v/1.2.0",
opts=ResourceOptions(depends_on=["aws_securityhub_account.example"]))
pci321 = aws.securityhub.StandardsSubscription("pci321", standards_arn="arn:aws:securityhub:us-east-1::standards/pci-dss/v/3.2.1",
opts=ResourceOptions(depends_on=["aws_securityhub_account.example"]))


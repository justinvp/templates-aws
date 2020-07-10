import pulumi
import pulumi_aws as aws

filter = aws.ses.ReceiptFilter("filter",
    cidr="10.10.10.10",
    policy="Block")


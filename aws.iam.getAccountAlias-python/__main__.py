import pulumi
import pulumi_aws as aws

current = aws.iam.get_account_alias()
pulumi.export("accountId", current.account_alias)


import pulumi
import pulumi_aws as aws

alias = aws.iam.AccountAlias("alias", account_alias="my-account-alias")


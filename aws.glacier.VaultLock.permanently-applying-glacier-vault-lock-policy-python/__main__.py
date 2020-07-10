import pulumi
import pulumi_aws as aws

example = aws.glacier.VaultLock("example",
    complete_lock=True,
    policy=data["aws_iam_policy_document"]["example"]["json"],
    vault_name=aws_glacier_vault["example"]["name"])


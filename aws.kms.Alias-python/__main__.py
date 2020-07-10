import pulumi
import pulumi_aws as aws

key = aws.kms.Key("key")
alias = aws.kms.Alias("alias", target_key_id=key.key_id)


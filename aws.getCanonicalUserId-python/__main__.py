import pulumi
import pulumi_aws as aws

current = aws.get_canonical_user_id()
pulumi.export("canonicalUserId", current.id)


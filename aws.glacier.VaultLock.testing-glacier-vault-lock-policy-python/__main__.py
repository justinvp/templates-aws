import pulumi
import pulumi_aws as aws

example_vault = aws.glacier.Vault("exampleVault")
example_policy_document = example_vault.arn.apply(lambda arn: aws.iam.get_policy_document(statements=[{
    "actions": ["glacier:DeleteArchive"],
    "condition": [{
        "test": "NumericLessThanEquals",
        "values": ["365"],
        "variable": "glacier:ArchiveAgeinDays",
    }],
    "effect": "Deny",
    "resources": [arn],
}]))
example_vault_lock = aws.glacier.VaultLock("exampleVaultLock",
    complete_lock=False,
    policy=example_policy_document.json,
    vault_name=example_vault.name)


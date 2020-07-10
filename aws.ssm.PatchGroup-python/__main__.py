import pulumi
import pulumi_aws as aws

production = aws.ssm.PatchBaseline("production", approved_patches=["KB123456"])
patchgroup = aws.ssm.PatchGroup("patchgroup",
    baseline_id=production.id,
    patch_group="patch-group-name")


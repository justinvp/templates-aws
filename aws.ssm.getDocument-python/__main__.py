import pulumi
import pulumi_aws as aws

foo = aws.ssm.get_document(document_format="YAML",
    name="AWS-GatherSoftwareInventory")
pulumi.export("content", foo.content)


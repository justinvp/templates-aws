import pulumi
import pulumi_aws as aws

foo = aws.ecr.Repository("foo",
    image_scanning_configuration={
        "scanOnPush": True,
    },
    image_tag_mutability="MUTABLE")


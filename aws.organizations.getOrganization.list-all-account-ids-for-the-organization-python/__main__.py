import pulumi
import pulumi_aws as aws

example = aws.organizations.get_organization()
pulumi.export("accountIds", [__item["id"] for __item in example.accounts])


import pulumi
import pulumi_aws as aws

example_resource_share = aws.ram.ResourceShare("exampleResourceShare", allow_external_principals=True)
example_principal_association = aws.ram.PrincipalAssociation("examplePrincipalAssociation",
    principal="111111111111",
    resource_share_arn=example_resource_share.arn)


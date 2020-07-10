import pulumi
import pulumi_aws as aws

org = aws.organizations.get_organization()
ou = aws.organizations.get_organizational_units(parent_id=org.roots[0]["id"])


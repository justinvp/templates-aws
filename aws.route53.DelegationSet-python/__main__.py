import pulumi
import pulumi_aws as aws

main = aws.route53.DelegationSet("main", reference_name="DynDNS")
primary = aws.route53.Zone("primary", delegation_set_id=main.id)
secondary = aws.route53.Zone("secondary", delegation_set_id=main.id)


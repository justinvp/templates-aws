import pulumi
import pulumi_aws as aws

ptfe_service_vpc_endpoint = aws.ec2.VpcEndpoint("ptfeServiceVpcEndpoint",
    private_dns_enabled=False,
    security_group_ids=[aws_security_group["ptfe_service"]["id"]],
    service_name=var["ptfe_service"],
    subnet_ids=[local["subnet_ids"]],
    vpc_endpoint_type="Interface",
    vpc_id=var["vpc_id"])
internal = aws.route53.get_zone(name="vpc.internal.",
    private_zone=True,
    vpc_id=var["vpc_id"])
ptfe_service_record = aws.route53.Record("ptfeServiceRecord",
    name=f"ptfe.{internal.name}",
    records=[ptfe_service_vpc_endpoint.dns_entries[0]["dns_name"]],
    ttl="300",
    type="CNAME",
    zone_id=internal.zone_id)


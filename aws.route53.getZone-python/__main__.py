import pulumi
import pulumi_aws as aws

selected = aws.route53.get_zone(name="test.com.",
    private_zone=True)
www = aws.route53.Record("www",
    name=f"www.{selected.name}",
    records=["10.0.0.1"],
    ttl="300",
    type="A",
    zone_id=selected.zone_id)


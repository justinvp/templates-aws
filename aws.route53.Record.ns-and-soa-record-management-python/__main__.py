import pulumi
import pulumi_aws as aws

example_zone = aws.route53.Zone("exampleZone")
example_record = aws.route53.Record("exampleRecord",
    allow_overwrite=True,
    name="test.example.com",
    records=[
        example_zone.name_servers[0],
        example_zone.name_servers[1],
        example_zone.name_servers[2],
        example_zone.name_servers[3],
    ],
    ttl=30,
    type="NS",
    zone_id=example_zone.zone_id)


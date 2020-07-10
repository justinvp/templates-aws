import pulumi
import pulumi_aws as aws

main = aws.route53.Zone("main")
dev = aws.route53.Zone("dev", tags={
    "Environment": "dev",
})
dev_ns = aws.route53.Record("dev-ns",
    name="dev.example.com",
    records=[
        dev.name_servers[0],
        dev.name_servers[1],
        dev.name_servers[2],
        dev.name_servers[3],
    ],
    ttl="30",
    type="NS",
    zone_id=main.zone_id)


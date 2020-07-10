import pulumi
import pulumi_aws as aws

www_dev = aws.route53.Record("www-dev",
    name="www",
    records=["dev.example.com"],
    set_identifier="dev",
    ttl="5",
    type="CNAME",
    weighted_routing_policies=[{
        "weight": 10,
    }],
    zone_id=aws_route53_zone["primary"]["zone_id"])
www_live = aws.route53.Record("www-live",
    name="www",
    records=["live.example.com"],
    set_identifier="live",
    ttl="5",
    type="CNAME",
    weighted_routing_policies=[{
        "weight": 90,
    }],
    zone_id=aws_route53_zone["primary"]["zone_id"])


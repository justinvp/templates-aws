import pulumi
import pulumi_aws as aws

www = aws.route53.Record("www",
    name="www.example.com",
    records=[aws_eip["lb"]["public_ip"]],
    ttl="300",
    type="A",
    zone_id=aws_route53_zone["primary"]["zone_id"])


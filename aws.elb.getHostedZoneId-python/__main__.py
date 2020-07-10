import pulumi
import pulumi_aws as aws

main = aws.elb.get_hosted_zone_id()
www = aws.route53.Record("www",
    aliases=[{
        "evaluateTargetHealth": True,
        "name": aws_elb["main"]["dns_name"],
        "zone_id": main.id,
    }],
    name="example.com",
    type="A",
    zone_id=aws_route53_zone["primary"]["zone_id"])


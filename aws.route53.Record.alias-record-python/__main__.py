import pulumi
import pulumi_aws as aws

main = aws.elb.LoadBalancer("main",
    availability_zones=["us-east-1c"],
    listeners=[{
        "instance_port": 80,
        "instanceProtocol": "http",
        "lb_port": 80,
        "lbProtocol": "http",
    }])
www = aws.route53.Record("www",
    aliases=[{
        "evaluateTargetHealth": True,
        "name": main.dns_name,
        "zone_id": main.zone_id,
    }],
    name="example.com",
    type="A",
    zone_id=aws_route53_zone["primary"]["zone_id"])


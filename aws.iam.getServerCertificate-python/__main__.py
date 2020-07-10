import pulumi
import pulumi_aws as aws

my_domain = aws.iam.get_server_certificate(latest=True,
    name_prefix="my-domain.org")
elb = aws.elb.LoadBalancer("elb", listeners=[{
    "instance_port": 8000,
    "instanceProtocol": "https",
    "lb_port": 443,
    "lbProtocol": "https",
    "sslCertificateId": my_domain.arn,
}])


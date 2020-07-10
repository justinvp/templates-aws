import pulumi
import pulumi_aws as aws

lb = aws.elb.LoadBalancer("lb",
    availability_zones=["us-east-1a"],
    listeners=[{
        "instance_port": 8000,
        "instanceProtocol": "https",
        "lb_port": 443,
        "lbProtocol": "https",
        "sslCertificateId": "arn:aws:iam::123456789012:server-certificate/certName",
    }])
foo = aws.elb.SslNegotiationPolicy("foo",
    attributes=[
        {
            "name": "Protocol-TLSv1",
            "value": "false",
        },
        {
            "name": "Protocol-TLSv1.1",
            "value": "false",
        },
        {
            "name": "Protocol-TLSv1.2",
            "value": "true",
        },
        {
            "name": "Server-Defined-Cipher-Order",
            "value": "true",
        },
        {
            "name": "ECDHE-RSA-AES128-GCM-SHA256",
            "value": "true",
        },
        {
            "name": "AES128-GCM-SHA256",
            "value": "true",
        },
        {
            "name": "EDH-RSA-DES-CBC3-SHA",
            "value": "false",
        },
    ],
    lb_port=443,
    load_balancer=lb.id)


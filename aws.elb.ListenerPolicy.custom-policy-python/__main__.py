import pulumi
import pulumi_aws as aws

wu_tang = aws.elb.LoadBalancer("wu-tang",
    availability_zones=["us-east-1a"],
    listeners=[{
        "instance_port": 443,
        "instanceProtocol": "http",
        "lb_port": 443,
        "lbProtocol": "https",
        "sslCertificateId": "arn:aws:iam::000000000000:server-certificate/wu-tang.net",
    }],
    tags={
        "Name": "wu-tang",
    })
wu_tang_ssl = aws.elb.LoadBalancerPolicy("wu-tang-ssl",
    load_balancer_name=wu_tang.name,
    policy_attributes=[
        {
            "name": "ECDHE-ECDSA-AES128-GCM-SHA256",
            "value": "true",
        },
        {
            "name": "Protocol-TLSv1.2",
            "value": "true",
        },
    ],
    policy_name="wu-tang-ssl",
    policy_type_name="SSLNegotiationPolicyType")
wu_tang_listener_policies_443 = aws.elb.ListenerPolicy("wu-tang-listener-policies-443",
    load_balancer_name=wu_tang.name,
    load_balancer_port=443,
    policy_names=[wu_tang_ssl.policy_name])


using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var wu_tang = new Aws.Elb.LoadBalancer("wu-tang", new Aws.Elb.LoadBalancerArgs
        {
            AvailabilityZones = 
            {
                "us-east-1a",
            },
            Listeners = 
            {
                new Aws.Elb.Inputs.LoadBalancerListenerArgs
                {
                    InstancePort = 443,
                    InstanceProtocol = "http",
                    LbPort = 443,
                    LbProtocol = "https",
                    SslCertificateId = "arn:aws:iam::000000000000:server-certificate/wu-tang.net",
                },
            },
            Tags = 
            {
                { "Name", "wu-tang" },
            },
        });
        var wu_tang_ssl = new Aws.Elb.LoadBalancerPolicy("wu-tang-ssl", new Aws.Elb.LoadBalancerPolicyArgs
        {
            LoadBalancerName = wu_tang.Name,
            PolicyAttributes = 
            {
                new Aws.Elb.Inputs.LoadBalancerPolicyPolicyAttributeArgs
                {
                    Name = "ECDHE-ECDSA-AES128-GCM-SHA256",
                    Value = "true",
                },
                new Aws.Elb.Inputs.LoadBalancerPolicyPolicyAttributeArgs
                {
                    Name = "Protocol-TLSv1.2",
                    Value = "true",
                },
            },
            PolicyName = "wu-tang-ssl",
            PolicyTypeName = "SSLNegotiationPolicyType",
        });
        var wu_tang_listener_policies_443 = new Aws.Elb.ListenerPolicy("wu-tang-listener-policies-443", new Aws.Elb.ListenerPolicyArgs
        {
            LoadBalancerName = wu_tang.Name,
            LoadBalancerPort = 443,
            PolicyNames = 
            {
                wu_tang_ssl.PolicyName,
            },
        });
    }

}


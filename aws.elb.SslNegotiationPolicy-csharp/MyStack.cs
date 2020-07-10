using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var lb = new Aws.Elb.LoadBalancer("lb", new Aws.Elb.LoadBalancerArgs
        {
            AvailabilityZones = 
            {
                "us-east-1a",
            },
            Listeners = 
            {
                new Aws.Elb.Inputs.LoadBalancerListenerArgs
                {
                    InstancePort = 8000,
                    InstanceProtocol = "https",
                    LbPort = 443,
                    LbProtocol = "https",
                    SslCertificateId = "arn:aws:iam::123456789012:server-certificate/certName",
                },
            },
        });
        var foo = new Aws.Elb.SslNegotiationPolicy("foo", new Aws.Elb.SslNegotiationPolicyArgs
        {
            Attributes = 
            {
                new Aws.Elb.Inputs.SslNegotiationPolicyAttributeArgs
                {
                    Name = "Protocol-TLSv1",
                    Value = "false",
                },
                new Aws.Elb.Inputs.SslNegotiationPolicyAttributeArgs
                {
                    Name = "Protocol-TLSv1.1",
                    Value = "false",
                },
                new Aws.Elb.Inputs.SslNegotiationPolicyAttributeArgs
                {
                    Name = "Protocol-TLSv1.2",
                    Value = "true",
                },
                new Aws.Elb.Inputs.SslNegotiationPolicyAttributeArgs
                {
                    Name = "Server-Defined-Cipher-Order",
                    Value = "true",
                },
                new Aws.Elb.Inputs.SslNegotiationPolicyAttributeArgs
                {
                    Name = "ECDHE-RSA-AES128-GCM-SHA256",
                    Value = "true",
                },
                new Aws.Elb.Inputs.SslNegotiationPolicyAttributeArgs
                {
                    Name = "AES128-GCM-SHA256",
                    Value = "true",
                },
                new Aws.Elb.Inputs.SslNegotiationPolicyAttributeArgs
                {
                    Name = "EDH-RSA-DES-CBC3-SHA",
                    Value = "false",
                },
            },
            LbPort = 443,
            LoadBalancer = lb.Id,
        });
    }

}


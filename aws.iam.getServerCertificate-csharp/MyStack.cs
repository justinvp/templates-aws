using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var my_domain = Output.Create(Aws.Iam.GetServerCertificate.InvokeAsync(new Aws.Iam.GetServerCertificateArgs
        {
            Latest = true,
            NamePrefix = "my-domain.org",
        }));
        var elb = new Aws.Elb.LoadBalancer("elb", new Aws.Elb.LoadBalancerArgs
        {
            Listeners = 
            {
                new Aws.Elb.Inputs.LoadBalancerListenerArgs
                {
                    InstancePort = 8000,
                    InstanceProtocol = "https",
                    LbPort = 443,
                    LbProtocol = "https",
                    SslCertificateId = my_domain.Apply(my_domain => my_domain.Arn),
                },
            },
        });
    }

}


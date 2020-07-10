using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleCertificate = new Aws.Acm.Certificate("exampleCertificate", new Aws.Acm.CertificateArgs
        {
        });
        var frontEndLoadBalancer = new Aws.LB.LoadBalancer("frontEndLoadBalancer", new Aws.LB.LoadBalancerArgs
        {
        });
        var frontEndListener = new Aws.LB.Listener("frontEndListener", new Aws.LB.ListenerArgs
        {
        });
        var exampleListenerCertificate = new Aws.LB.ListenerCertificate("exampleListenerCertificate", new Aws.LB.ListenerCertificateArgs
        {
            CertificateArn = exampleCertificate.Arn,
            ListenerArn = frontEndListener.Arn,
        });
    }

}


using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Route53.HealthCheck("example", new Aws.Route53.HealthCheckArgs
        {
            FailureThreshold = 5,
            Fqdn = "example.com",
            Port = 443,
            RequestInterval = 30,
            ResourcePath = "/",
            SearchString = "example",
            Type = "HTTPS_STR_MATCH",
        });
    }

}


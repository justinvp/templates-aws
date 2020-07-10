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
            Port = 80,
            RequestInterval = 30,
            ResourcePath = "/",
            Tags = 
            {
                { "Name", "tf-test-health-check" },
            },
            Type = "HTTP",
        });
    }

}


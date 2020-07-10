using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var parent = new Aws.Route53.HealthCheck("parent", new Aws.Route53.HealthCheckArgs
        {
            ChildHealthThreshold = 1,
            ChildHealthchecks = 
            {
                aws_route53_health_check.Child.Id,
            },
            Tags = 
            {
                { "Name", "tf-test-calculated-health-check" },
            },
            Type = "CALCULATED",
        });
    }

}


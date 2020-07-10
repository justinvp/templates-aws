using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        // Create a new load balancer
        var bar = new Aws.Elb.LoadBalancer("bar", new Aws.Elb.LoadBalancerArgs
        {
            AccessLogs = new Aws.Elb.Inputs.LoadBalancerAccessLogsArgs
            {
                Bucket = "foo",
                BucketPrefix = "bar",
                Interval = 60,
            },
            AvailabilityZones = 
            {
                "us-west-2a",
                "us-west-2b",
                "us-west-2c",
            },
            ConnectionDraining = true,
            ConnectionDrainingTimeout = 400,
            CrossZoneLoadBalancing = true,
            HealthCheck = new Aws.Elb.Inputs.LoadBalancerHealthCheckArgs
            {
                HealthyThreshold = 2,
                Interval = 30,
                Target = "HTTP:8000/",
                Timeout = 3,
                UnhealthyThreshold = 2,
            },
            IdleTimeout = 400,
            Instances = 
            {
                aws_instance.Foo.Id,
            },
            Listeners = 
            {
                new Aws.Elb.Inputs.LoadBalancerListenerArgs
                {
                    InstancePort = 8000,
                    InstanceProtocol = "http",
                    LbPort = 80,
                    LbProtocol = "http",
                },
                new Aws.Elb.Inputs.LoadBalancerListenerArgs
                {
                    InstancePort = 8000,
                    InstanceProtocol = "http",
                    LbPort = 443,
                    LbProtocol = "https",
                    SslCertificateId = "arn:aws:iam::123456789012:server-certificate/certName",
                },
            },
            Tags = 
            {
                { "Name", "foobar-elb" },
            },
        });
    }

}


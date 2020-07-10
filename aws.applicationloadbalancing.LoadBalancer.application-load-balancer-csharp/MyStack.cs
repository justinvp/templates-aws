using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var test = new Aws.LB.LoadBalancer("test", new Aws.LB.LoadBalancerArgs
        {
            AccessLogs = new Aws.LB.Inputs.LoadBalancerAccessLogsArgs
            {
                Bucket = aws_s3_bucket.Lb_logs.Bucket,
                Enabled = true,
                Prefix = "test-lb",
            },
            EnableDeletionProtection = true,
            Internal = false,
            LoadBalancerType = "application",
            SecurityGroups = 
            {
                aws_security_group.Lb_sg.Id,
            },
            Subnets = 
            {
                aws_subnet.Public.Select(__item => __item.Id).ToList(),
            },
            Tags = 
            {
                { "Environment", "production" },
            },
        });
    }

}


using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var main = Output.Create(Aws.Elb.GetServiceAccount.InvokeAsync());
        var elbLogs = new Aws.S3.Bucket("elbLogs", new Aws.S3.BucketArgs
        {
            Acl = "private",
            Policy = main.Apply(main => @$"{{
  ""Id"": ""Policy"",
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {{
      ""Action"": [
        ""s3:PutObject""
      ],
      ""Effect"": ""Allow"",
      ""Resource"": ""arn:aws:s3:::my-elb-tf-test-bucket/AWSLogs/*"",
      ""Principal"": {{
        ""AWS"": [
          ""{main.Arn}""
        ]
      }}
    }}
  ]
}}

"),
        });
        var bar = new Aws.Elb.LoadBalancer("bar", new Aws.Elb.LoadBalancerArgs
        {
            AccessLogs = new Aws.Elb.Inputs.LoadBalancerAccessLogsArgs
            {
                Bucket = elbLogs.BucketName,
                Interval = 5,
            },
            AvailabilityZones = 
            {
                "us-west-2a",
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
            },
        });
    }

}


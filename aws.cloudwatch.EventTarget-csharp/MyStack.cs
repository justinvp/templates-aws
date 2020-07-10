using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var console = new Aws.CloudWatch.EventRule("console", new Aws.CloudWatch.EventRuleArgs
        {
            Description = "Capture all EC2 scaling events",
            EventPattern = @"{
  ""source"": [
    ""aws.autoscaling""
  ],
  ""detail-type"": [
    ""EC2 Instance Launch Successful"",
    ""EC2 Instance Terminate Successful"",
    ""EC2 Instance Launch Unsuccessful"",
    ""EC2 Instance Terminate Unsuccessful""
  ]
}

",
        });
        var testStream = new Aws.Kinesis.Stream("testStream", new Aws.Kinesis.StreamArgs
        {
            ShardCount = 1,
        });
        var yada = new Aws.CloudWatch.EventTarget("yada", new Aws.CloudWatch.EventTargetArgs
        {
            Arn = testStream.Arn,
            Rule = console.Name,
            RunCommandTargets = 
            {
                new Aws.CloudWatch.Inputs.EventTargetRunCommandTargetArgs
                {
                    Key = "tag:Name",
                    Values = 
                    {
                        "FooBar",
                    },
                },
                new Aws.CloudWatch.Inputs.EventTargetRunCommandTargetArgs
                {
                    Key = "InstanceIds",
                    Values = 
                    {
                        "i-162058cd308bffec2",
                    },
                },
            },
        });
    }

}


using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleLogGroup = new Aws.CloudWatch.LogGroup("exampleLogGroup", new Aws.CloudWatch.LogGroupArgs
        {
        });
        var exampleLogResourcePolicy = new Aws.CloudWatch.LogResourcePolicy("exampleLogResourcePolicy", new Aws.CloudWatch.LogResourcePolicyArgs
        {
            PolicyDocument = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {
      ""Effect"": ""Allow"",
      ""Principal"": {
        ""Service"": ""es.amazonaws.com""
      },
      ""Action"": [
        ""logs:PutLogEvents"",
        ""logs:PutLogEventsBatch"",
        ""logs:CreateLogStream""
      ],
      ""Resource"": ""arn:aws:logs:*""
    }
  ]
}

",
            PolicyName = "example",
        });
        var exampleDomain = new Aws.ElasticSearch.Domain("exampleDomain", new Aws.ElasticSearch.DomainArgs
        {
            LogPublishingOptions = 
            {
                new Aws.ElasticSearch.Inputs.DomainLogPublishingOptionArgs
                {
                    CloudwatchLogGroupArn = exampleLogGroup.Arn,
                    LogType = "INDEX_SLOW_LOGS",
                },
            },
        });
    }

}


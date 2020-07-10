using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var queue = new Aws.Sqs.Queue("queue", new Aws.Sqs.QueueArgs
        {
        });
        var test = new Aws.Sqs.QueuePolicy("test", new Aws.Sqs.QueuePolicyArgs
        {
            Policy = queue.Arn.Apply(arn => @$"{{
  ""Version"": ""2012-10-17"",
  ""Id"": ""sqspolicy"",
  ""Statement"": [
    {{
      ""Sid"": ""First"",
      ""Effect"": ""Allow"",
      ""Principal"": ""*"",
      ""Action"": ""sqs:SendMessage"",
      ""Resource"": ""{arn}"",
      ""Condition"": {{
        ""ArnEquals"": {{
          ""aws:SourceArn"": ""{aws_sns_topic.Example.Arn}""
        }}
      }}
    }}
  ]
}}

"),
            QueueUrl = queue.Id,
        });
    }

}


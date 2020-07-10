using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var app = new Aws.Pinpoint.App("app", new Aws.Pinpoint.AppArgs
        {
        });
        var testStream = new Aws.Kinesis.Stream("testStream", new Aws.Kinesis.StreamArgs
        {
            ShardCount = 1,
        });
        var testRole = new Aws.Iam.Role("testRole", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {
      ""Action"": ""sts:AssumeRole"",
      ""Principal"": {
        ""Service"": ""pinpoint.us-east-1.amazonaws.com""
      },
      ""Effect"": ""Allow"",
      ""Sid"": """"
    }
  ]
}

",
        });
        var stream = new Aws.Pinpoint.EventStream("stream", new Aws.Pinpoint.EventStreamArgs
        {
            ApplicationId = app.ApplicationId,
            DestinationStreamArn = testStream.Arn,
            RoleArn = testRole.Arn,
        });
        var testRolePolicy = new Aws.Iam.RolePolicy("testRolePolicy", new Aws.Iam.RolePolicyArgs
        {
            Policy = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": {
    ""Action"": [
      ""kinesis:PutRecords"",
      ""kinesis:DescribeStream""
    ],
    ""Effect"": ""Allow"",
    ""Resource"": [
      ""arn:aws:kinesis:us-east-1:*:*/*""
    ]
  }
}

",
            Role = testRole.Id,
        });
    }

}


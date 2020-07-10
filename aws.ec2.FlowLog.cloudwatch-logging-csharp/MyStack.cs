using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleLogGroup = new Aws.CloudWatch.LogGroup("exampleLogGroup", new Aws.CloudWatch.LogGroupArgs
        {
        });
        var exampleRole = new Aws.Iam.Role("exampleRole", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {
      ""Sid"": """",
      ""Effect"": ""Allow"",
      ""Principal"": {
        ""Service"": ""vpc-flow-logs.amazonaws.com""
      },
      ""Action"": ""sts:AssumeRole""
    }
  ]
}

",
        });
        var exampleFlowLog = new Aws.Ec2.FlowLog("exampleFlowLog", new Aws.Ec2.FlowLogArgs
        {
            IamRoleArn = exampleRole.Arn,
            LogDestination = exampleLogGroup.Arn,
            TrafficType = "ALL",
            VpcId = aws_vpc.Example.Id,
        });
        var exampleRolePolicy = new Aws.Iam.RolePolicy("exampleRolePolicy", new Aws.Iam.RolePolicyArgs
        {
            Policy = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {
      ""Action"": [
        ""logs:CreateLogGroup"",
        ""logs:CreateLogStream"",
        ""logs:PutLogEvents"",
        ""logs:DescribeLogGroups"",
        ""logs:DescribeLogStreams""
      ],
      ""Effect"": ""Allow"",
      ""Resource"": ""*""
    }
  ]
}

",
            Role = exampleRole.Id,
        });
    }

}


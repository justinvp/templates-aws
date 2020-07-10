using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var sfnStateMachine = new Aws.Sfn.StateMachine("sfnStateMachine", new Aws.Sfn.StateMachineArgs
        {
            Definition = @$"{{
  ""Comment"": ""A Hello World example of the Amazon States Language using an AWS Lambda Function"",
  ""StartAt"": ""HelloWorld"",
  ""States"": {{
    ""HelloWorld"": {{
      ""Type"": ""Task"",
      ""Resource"": ""{aws_lambda_function.Lambda.Arn}"",
      ""End"": true
    }}
  }}
}}

",
            RoleArn = aws_iam_role.Iam_for_sfn.Arn,
        });
    }

}


using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Lambda.FunctionEventInvokeConfig("example", new Aws.Lambda.FunctionEventInvokeConfigArgs
        {
            FunctionName = aws_lambda_alias.Example.Function_name,
            MaximumEventAgeInSeconds = 60,
            MaximumRetryAttempts = 0,
        });
    }

}


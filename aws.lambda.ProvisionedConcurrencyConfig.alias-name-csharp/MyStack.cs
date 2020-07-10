using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Lambda.ProvisionedConcurrencyConfig("example", new Aws.Lambda.ProvisionedConcurrencyConfigArgs
        {
            FunctionName = aws_lambda_alias.Example.Function_name,
            ProvisionedConcurrentExecutions = 1,
            Qualifier = aws_lambda_alias.Example.Name,
        });
    }

}


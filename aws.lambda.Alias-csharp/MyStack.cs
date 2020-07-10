using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var testLambdaAlias = new Aws.Lambda.Alias("testLambdaAlias", new Aws.Lambda.AliasArgs
        {
            Description = "a sample description",
            FunctionName = aws_lambda_function.Lambda_function_test.Arn,
            FunctionVersion = "1",
            RoutingConfig = new Aws.Lambda.Inputs.AliasRoutingConfigArgs
            {
                AdditionalVersionWeights = 
                {
                    { "2", 0.5 },
                },
            },
        });
    }

}


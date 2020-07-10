using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var config = new Config();
        var functionName = config.RequireObject<dynamic>("functionName");
        var existing = Output.Create(Aws.Lambda.GetFunction.InvokeAsync(new Aws.Lambda.GetFunctionArgs
        {
            FunctionName = functionName,
        }));
    }

}


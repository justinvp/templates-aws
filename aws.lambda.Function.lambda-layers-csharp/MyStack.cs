using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleLayerVersion = new Aws.Lambda.LayerVersion("exampleLayerVersion", new Aws.Lambda.LayerVersionArgs
        {
        });
        var exampleFunction = new Aws.Lambda.Function("exampleFunction", new Aws.Lambda.FunctionArgs
        {
            Layers = 
            {
                exampleLayerVersion.Arn,
            },
        });
    }

}


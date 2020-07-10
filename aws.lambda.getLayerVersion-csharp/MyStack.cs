using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var config = new Config();
        var layerName = config.RequireObject<dynamic>("layerName");
        var existing = Output.Create(Aws.Lambda.GetLayerVersion.InvokeAsync(new Aws.Lambda.GetLayerVersionArgs
        {
            LayerName = layerName,
        }));
    }

}


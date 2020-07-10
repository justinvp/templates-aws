using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var dada = new Aws.CloudWatch.LogGroup("dada", new Aws.CloudWatch.LogGroupArgs
        {
        });
        var yada = new Aws.CloudWatch.LogMetricFilter("yada", new Aws.CloudWatch.LogMetricFilterArgs
        {
            LogGroupName = dada.Name,
            MetricTransformation = new Aws.CloudWatch.Inputs.LogMetricFilterMetricTransformationArgs
            {
                Name = "EventCount",
                Namespace = "YourNamespace",
                Value = "1",
            },
            Pattern = "",
        });
    }

}


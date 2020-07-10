using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var reportDefinition = Output.Create(Aws.Cur.GetReportDefinition.InvokeAsync(new Aws.Cur.GetReportDefinitionArgs
        {
            ReportName = "example",
        }));
    }

}


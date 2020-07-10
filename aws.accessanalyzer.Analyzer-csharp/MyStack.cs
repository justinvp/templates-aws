using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.AccessAnalyzer.Analyzer("example", new Aws.AccessAnalyzer.AnalyzerArgs
        {
            AnalyzerName = "example",
        });
    }

}


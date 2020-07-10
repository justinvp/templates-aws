using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Glue.Classifier("example", new Aws.Glue.ClassifierArgs
        {
            GrokClassifier = new Aws.Glue.Inputs.ClassifierGrokClassifierArgs
            {
                Classification = "example",
                GrokPattern = "example",
            },
        });
    }

}


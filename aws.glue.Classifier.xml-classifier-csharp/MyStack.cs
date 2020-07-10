using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Glue.Classifier("example", new Aws.Glue.ClassifierArgs
        {
            XmlClassifier = new Aws.Glue.Inputs.ClassifierXmlClassifierArgs
            {
                Classification = "example",
                RowTag = "example",
            },
        });
    }

}


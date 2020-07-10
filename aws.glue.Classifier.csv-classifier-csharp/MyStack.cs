using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Glue.Classifier("example", new Aws.Glue.ClassifierArgs
        {
            CsvClassifier = new Aws.Glue.Inputs.ClassifierCsvClassifierArgs
            {
                AllowSingleColumn = false,
                ContainsHeader = "PRESENT",
                Delimiter = ",",
                DisableValueTrimming = false,
                Header = 
                {
                    "example1",
                    "example2",
                },
                QuoteSymbol = "'",
            },
        });
    }

}


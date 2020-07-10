using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var current = Output.Create(Aws.GetPartition.InvokeAsync());
        var s3Policy = current.Apply(current => Output.Create(Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
        {
            Statements = 
            {
                new Aws.Iam.Inputs.GetPolicyDocumentStatementArgs
                {
                    Actions = 
                    {
                        "s3:ListBucket",
                    },
                    Resources = 
                    {
                        $"arn:{current.Partition}:s3:::my-bucket",
                    },
                    Sid = "1",
                },
            },
        })));
    }

}


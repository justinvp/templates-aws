using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Macie.S3BucketAssociation("example", new Aws.Macie.S3BucketAssociationArgs
        {
            BucketName = "tf-macie-example",
            ClassificationType = new Aws.Macie.Inputs.S3BucketAssociationClassificationTypeArgs
            {
                OneTime = "FULL",
            },
            Prefix = "data",
        });
    }

}


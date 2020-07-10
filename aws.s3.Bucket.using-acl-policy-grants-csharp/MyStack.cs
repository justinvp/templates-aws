using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var currentUser = Output.Create(Aws.GetCanonicalUserId.InvokeAsync());
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {
            Grants = 
            {
                new Aws.S3.Inputs.BucketGrantArgs
                {
                    Id = currentUser.Apply(currentUser => currentUser.Id),
                    Permissions = 
                    {
                        "FULL_CONTROL",
                    },
                    Type = "CanonicalUser",
                },
                new Aws.S3.Inputs.BucketGrantArgs
                {
                    Permissions = 
                    {
                        "READ",
                        "WRITE",
                    },
                    Type = "Group",
                    Uri = "http://acs.amazonaws.com/groups/s3/LogDelivery",
                },
            },
        });
    }

}


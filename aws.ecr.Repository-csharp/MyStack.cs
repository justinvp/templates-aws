using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var foo = new Aws.Ecr.Repository("foo", new Aws.Ecr.RepositoryArgs
        {
            ImageScanningConfiguration = new Aws.Ecr.Inputs.RepositoryImageScanningConfigurationArgs
            {
                ScanOnPush = true,
            },
            ImageTagMutability = "MUTABLE",
        });
    }

}


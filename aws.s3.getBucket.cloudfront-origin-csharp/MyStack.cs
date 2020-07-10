using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var selected = Output.Create(Aws.S3.GetBucket.InvokeAsync(new Aws.S3.GetBucketArgs
        {
            Bucket = "a-test-bucket",
        }));
        var test = new Aws.CloudFront.Distribution("test", new Aws.CloudFront.DistributionArgs
        {
            Origins = 
            {
                new Aws.CloudFront.Inputs.DistributionOriginArgs
                {
                    DomainName = selected.Apply(selected => selected.BucketDomainName),
                    OriginId = "s3-selected-bucket",
                },
            },
        });
    }

}


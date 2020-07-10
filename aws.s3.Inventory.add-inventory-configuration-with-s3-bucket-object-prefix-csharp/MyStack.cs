using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var test = new Aws.S3.Bucket("test", new Aws.S3.BucketArgs
        {
        });
        var inventory = new Aws.S3.Bucket("inventory", new Aws.S3.BucketArgs
        {
        });
        var test_prefix = new Aws.S3.Inventory("test-prefix", new Aws.S3.InventoryArgs
        {
            Bucket = test.Id,
            Destination = new Aws.S3.Inputs.InventoryDestinationArgs
            {
                Bucket = new Aws.S3.Inputs.InventoryDestinationBucketArgs
                {
                    BucketArn = inventory.Arn,
                    Format = "ORC",
                    Prefix = "inventory",
                },
            },
            Filter = new Aws.S3.Inputs.InventoryFilterArgs
            {
                Prefix = "documents/",
            },
            IncludedObjectVersions = "All",
            Schedule = new Aws.S3.Inputs.InventoryScheduleArgs
            {
                Frequency = "Daily",
            },
        });
    }

}


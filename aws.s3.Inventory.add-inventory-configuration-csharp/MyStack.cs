using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var testBucket = new Aws.S3.Bucket("testBucket", new Aws.S3.BucketArgs
        {
        });
        var inventory = new Aws.S3.Bucket("inventory", new Aws.S3.BucketArgs
        {
        });
        var testInventory = new Aws.S3.Inventory("testInventory", new Aws.S3.InventoryArgs
        {
            Bucket = testBucket.Id,
            Destination = new Aws.S3.Inputs.InventoryDestinationArgs
            {
                Bucket = new Aws.S3.Inputs.InventoryDestinationBucketArgs
                {
                    BucketArn = inventory.Arn,
                    Format = "ORC",
                },
            },
            IncludedObjectVersions = "All",
            Schedule = new Aws.S3.Inputs.InventoryScheduleArgs
            {
                Frequency = "Daily",
            },
        });
    }

}


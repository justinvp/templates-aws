using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var hogeBucket = new Aws.S3.Bucket("hogeBucket", new Aws.S3.BucketArgs
        {
        });
        var testKey = new Aws.Kms.Key("testKey", new Aws.Kms.KeyArgs
        {
            DeletionWindowInDays = 7,
            Description = "Athena KMS Key",
        });
        var testWorkgroup = new Aws.Athena.Workgroup("testWorkgroup", new Aws.Athena.WorkgroupArgs
        {
            Configuration = new Aws.Athena.Inputs.WorkgroupConfigurationArgs
            {
                ResultConfiguration = new Aws.Athena.Inputs.WorkgroupConfigurationResultConfigurationArgs
                {
                    EncryptionConfiguration = new Aws.Athena.Inputs.WorkgroupConfigurationResultConfigurationEncryptionConfigurationArgs
                    {
                        EncryptionOption = "SSE_KMS",
                        KmsKeyArn = testKey.Arn,
                    },
                },
            },
        });
        var hogeDatabase = new Aws.Athena.Database("hogeDatabase", new Aws.Athena.DatabaseArgs
        {
            Bucket = hogeBucket.Id,
            Name = "users",
        });
        var foo = new Aws.Athena.NamedQuery("foo", new Aws.Athena.NamedQueryArgs
        {
            Database = hogeDatabase.Name,
            Query = hogeDatabase.Name.Apply(name => $"SELECT * FROM {name} limit 10;"),
            Workgroup = testWorkgroup.Id,
        });
    }

}


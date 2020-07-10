using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var awsSnsTopic = new Aws.Sns.Topic("awsSnsTopic", new Aws.Sns.TopicArgs
        {
        });
        var myArchive = new Aws.Glacier.Vault("myArchive", new Aws.Glacier.VaultArgs
        {
            AccessPolicy = @"{
    ""Version"":""2012-10-17"",
    ""Statement"":[
       {
          ""Sid"": ""add-read-only-perm"",
          ""Principal"": ""*"",
          ""Effect"": ""Allow"",
          ""Action"": [
             ""glacier:InitiateJob"",
             ""glacier:GetJobOutput""
          ],
          ""Resource"": ""arn:aws:glacier:eu-west-1:432981146916:vaults/MyArchive""
       }
    ]
}

",
            Notifications = 
            {
                new Aws.Glacier.Inputs.VaultNotificationArgs
                {
                    Events = 
                    {
                        "ArchiveRetrievalCompleted",
                        "InventoryRetrievalCompleted",
                    },
                    SnsTopic = awsSnsTopic.Arn,
                },
            },
            Tags = 
            {
                { "Test", "MyArchive" },
            },
        });
    }

}


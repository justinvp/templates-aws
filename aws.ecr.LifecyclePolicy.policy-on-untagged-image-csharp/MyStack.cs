using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var foo = new Aws.Ecr.Repository("foo", new Aws.Ecr.RepositoryArgs
        {
        });
        var foopolicy = new Aws.Ecr.LifecyclePolicy("foopolicy", new Aws.Ecr.LifecyclePolicyArgs
        {
            Policy = @"{
    ""rules"": [
        {
            ""rulePriority"": 1,
            ""description"": ""Expire images older than 14 days"",
            ""selection"": {
                ""tagStatus"": ""untagged"",
                ""countType"": ""sinceImagePushed"",
                ""countUnit"": ""days"",
                ""countNumber"": 14
            },
            ""action"": {
                ""type"": ""expire""
            }
        }
    ]
}

",
            Repository = foo.Name,
        });
    }

}


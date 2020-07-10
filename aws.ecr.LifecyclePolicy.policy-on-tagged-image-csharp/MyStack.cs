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
            ""description"": ""Keep last 30 images"",
            ""selection"": {
                ""tagStatus"": ""tagged"",
                ""tagPrefixList"": [""v""],
                ""countType"": ""imageCountMoreThan"",
                ""countNumber"": 30
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


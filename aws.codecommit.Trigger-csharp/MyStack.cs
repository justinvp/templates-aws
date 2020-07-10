using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var testRepository = new Aws.CodeCommit.Repository("testRepository", new Aws.CodeCommit.RepositoryArgs
        {
            RepositoryName = "test",
        });
        var testTrigger = new Aws.CodeCommit.Trigger("testTrigger", new Aws.CodeCommit.TriggerArgs
        {
            RepositoryName = testRepository.RepositoryName,
            Triggers = 
            {
                new Aws.CodeCommit.Inputs.TriggerTriggerArgs
                {
                    DestinationArn = aws_sns_topic.Test.Arn,
                    Events = 
                    {
                        "all",
                    },
                    Name = "all",
                },
            },
        });
    }

}


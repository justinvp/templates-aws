using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var elasticbeanstalk = new Aws.Iam.ServiceLinkedRole("elasticbeanstalk", new Aws.Iam.ServiceLinkedRoleArgs
        {
            AwsServiceName = "elasticbeanstalk.amazonaws.com",
        });
    }

}


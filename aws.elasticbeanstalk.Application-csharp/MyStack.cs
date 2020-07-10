using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var tftest = new Aws.ElasticBeanstalk.Application("tftest", new Aws.ElasticBeanstalk.ApplicationArgs
        {
            AppversionLifecycle = new Aws.ElasticBeanstalk.Inputs.ApplicationAppversionLifecycleArgs
            {
                DeleteSourceFromS3 = true,
                MaxCount = 128,
                ServiceRole = aws_iam_role.Beanstalk_service.Arn,
            },
            Description = "tf-test-desc",
        });
    }

}


using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var tftest = new Aws.ElasticBeanstalk.Application("tftest", new Aws.ElasticBeanstalk.ApplicationArgs
        {
            Description = "tf-test-desc",
        });
        var tfenvtest = new Aws.ElasticBeanstalk.Environment("tfenvtest", new Aws.ElasticBeanstalk.EnvironmentArgs
        {
            Application = tftest.Name,
            SolutionStackName = "64bit Amazon Linux 2015.03 v2.0.3 running Go 1.4",
        });
    }

}


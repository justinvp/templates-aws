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
        var tfTemplate = new Aws.ElasticBeanstalk.ConfigurationTemplate("tfTemplate", new Aws.ElasticBeanstalk.ConfigurationTemplateArgs
        {
            Application = tftest.Name,
            SolutionStackName = "64bit Amazon Linux 2015.09 v2.0.8 running Go 1.4",
        });
    }

}


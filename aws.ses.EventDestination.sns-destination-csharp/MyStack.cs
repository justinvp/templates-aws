using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var sns = new Aws.Ses.EventDestination("sns", new Aws.Ses.EventDestinationArgs
        {
            ConfigurationSetName = aws_ses_configuration_set.Example.Name,
            Enabled = true,
            MatchingTypes = 
            {
                "bounce",
                "send",
            },
            SnsDestination = new Aws.Ses.Inputs.EventDestinationSnsDestinationArgs
            {
                TopicArn = aws_sns_topic.Example.Arn,
            },
        });
    }

}


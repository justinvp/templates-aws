using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var cloudwatch = new Aws.Ses.EventDestination("cloudwatch", new Aws.Ses.EventDestinationArgs
        {
            CloudwatchDestinations = 
            {
                new Aws.Ses.Inputs.EventDestinationCloudwatchDestinationArgs
                {
                    DefaultValue = "default",
                    DimensionName = "dimension",
                    ValueSource = "emailHeader",
                },
            },
            ConfigurationSetName = aws_ses_configuration_set.Example.Name,
            Enabled = true,
            MatchingTypes = 
            {
                "bounce",
                "send",
            },
        });
    }

}


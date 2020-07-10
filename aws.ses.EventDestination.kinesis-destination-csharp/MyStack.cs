using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var kinesis = new Aws.Ses.EventDestination("kinesis", new Aws.Ses.EventDestinationArgs
        {
            ConfigurationSetName = aws_ses_configuration_set.Example.Name,
            Enabled = true,
            KinesisDestination = new Aws.Ses.Inputs.EventDestinationKinesisDestinationArgs
            {
                RoleArn = aws_iam_role.Example.Arn,
                StreamArn = aws_kinesis_firehose_delivery_stream.Example.Arn,
            },
            MatchingTypes = 
            {
                "bounce",
                "send",
            },
        });
    }

}


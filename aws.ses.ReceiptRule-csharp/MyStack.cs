using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        // Add a header to the email and store it in S3
        var store = new Aws.Ses.ReceiptRule("store", new Aws.Ses.ReceiptRuleArgs
        {
            AddHeaderActions = 
            {
                new Aws.Ses.Inputs.ReceiptRuleAddHeaderActionArgs
                {
                    HeaderName = "Custom-Header",
                    HeaderValue = "Added by SES",
                    Position = 1,
                },
            },
            Enabled = true,
            Recipients = 
            {
                "karen@example.com",
            },
            RuleSetName = "default-rule-set",
            S3Actions = 
            {
                new Aws.Ses.Inputs.ReceiptRuleS3ActionArgs
                {
                    BucketName = "emails",
                    Position = 2,
                },
            },
            ScanEnabled = true,
        });
    }

}


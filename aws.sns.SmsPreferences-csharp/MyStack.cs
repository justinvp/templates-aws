using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var updateSmsPrefs = new Aws.Sns.SmsPreferences("updateSmsPrefs", new Aws.Sns.SmsPreferencesArgs
        {
        });
    }

}


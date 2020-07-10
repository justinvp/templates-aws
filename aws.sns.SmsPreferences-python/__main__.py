import pulumi
import pulumi_aws as aws

update_sms_prefs = aws.sns.SmsPreferences("updateSmsPrefs")


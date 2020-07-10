import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const updateSmsPrefs = new aws.sns.SmsPreferences("update_sms_prefs", {});


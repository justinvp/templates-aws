import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const awsSnsTopic = new aws.sns.Topic("aws_sns_topic", {});
const myArchive = new aws.glacier.Vault("my_archive", {
    accessPolicy: `{
    "Version":"2012-10-17",
    "Statement":[
       {
          "Sid": "add-read-only-perm",
          "Principal": "*",
          "Effect": "Allow",
          "Action": [
             "glacier:InitiateJob",
             "glacier:GetJobOutput"
          ],
          "Resource": "arn:aws:glacier:eu-west-1:432981146916:vaults/MyArchive"
       }
    ]
}
`,
    notifications: [{
        events: [
            "ArchiveRetrievalCompleted",
            "InventoryRetrievalCompleted",
        ],
        snsTopic: awsSnsTopic.arn,
    }],
    tags: {
        Test: "MyArchive",
    },
});


import pulumi
import pulumi_aws as aws

aws_sns_topic = aws.sns.Topic("awsSnsTopic")
my_archive = aws.glacier.Vault("myArchive",
    access_policy="""{
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

""",
    notifications=[{
        "events": [
            "ArchiveRetrievalCompleted",
            "InventoryRetrievalCompleted",
        ],
        "sns_topic": aws_sns_topic.arn,
    }],
    tags={
        "Test": "MyArchive",
    })


import pulumi
import pulumi_aws as aws

dlm_lifecycle_role = aws.iam.Role("dlmLifecycleRole", assume_role_policy="""{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "dlm.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}

""")
dlm_lifecycle = aws.iam.RolePolicy("dlmLifecycle",
    policy="""{
   "Version": "2012-10-17",
   "Statement": [
      {
         "Effect": "Allow",
         "Action": [
            "ec2:CreateSnapshot",
            "ec2:DeleteSnapshot",
            "ec2:DescribeVolumes",
            "ec2:DescribeSnapshots"
         ],
         "Resource": "*"
      },
      {
         "Effect": "Allow",
         "Action": [
            "ec2:CreateTags"
         ],
         "Resource": "arn:aws:ec2:*::snapshot/*"
      }
   ]
}

""",
    role=dlm_lifecycle_role.id)
example = aws.dlm.LifecyclePolicy("example",
    description="example DLM lifecycle policy",
    execution_role_arn=dlm_lifecycle_role.arn,
    policy_details={
        "resourceTypes": ["VOLUME"],
        "schedule": [{
            "copyTags": False,
            "createRule": {
                "interval": 24,
                "intervalUnit": "HOURS",
                "times": "23:45",
            },
            "name": "2 weeks of daily snapshots",
            "retainRule": {
                "count": 14,
            },
            "tagsToAdd": {
                "SnapshotCreator": "DLM",
            },
        }],
        "targetTags": {
            "Snapshot": "true",
        },
    },
    state="ENABLED")


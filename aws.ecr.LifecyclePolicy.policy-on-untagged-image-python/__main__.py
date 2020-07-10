import pulumi
import pulumi_aws as aws

foo = aws.ecr.Repository("foo")
foopolicy = aws.ecr.LifecyclePolicy("foopolicy",
    policy="""{
    "rules": [
        {
            "rulePriority": 1,
            "description": "Expire images older than 14 days",
            "selection": {
                "tagStatus": "untagged",
                "countType": "sinceImagePushed",
                "countUnit": "days",
                "countNumber": 14
            },
            "action": {
                "type": "expire"
            }
        }
    ]
}

""",
    repository=foo.name)


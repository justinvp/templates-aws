import pulumi
import pulumi_aws as aws

foo = aws.ecr.Repository("foo")
foopolicy = aws.ecr.LifecyclePolicy("foopolicy",
    policy="""{
    "rules": [
        {
            "rulePriority": 1,
            "description": "Keep last 30 images",
            "selection": {
                "tagStatus": "tagged",
                "tagPrefixList": ["v"],
                "countType": "imageCountMoreThan",
                "countNumber": 30
            },
            "action": {
                "type": "expire"
            }
        }
    ]
}

""",
    repository=foo.name)


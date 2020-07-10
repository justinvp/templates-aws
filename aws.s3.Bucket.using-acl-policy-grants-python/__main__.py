import pulumi
import pulumi_aws as aws

current_user = aws.get_canonical_user_id()
bucket = aws.s3.Bucket("bucket", grants=[
    {
        "id": current_user.id,
        "permissions": ["FULL_CONTROL"],
        "type": "CanonicalUser",
    },
    {
        "permissions": [
            "READ",
            "WRITE",
        ],
        "type": "Group",
        "uri": "http://acs.amazonaws.com/groups/s3/LogDelivery",
    },
])


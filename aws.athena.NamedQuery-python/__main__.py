import pulumi
import pulumi_aws as aws

hoge_bucket = aws.s3.Bucket("hogeBucket")
test_key = aws.kms.Key("testKey",
    deletion_window_in_days=7,
    description="Athena KMS Key")
test_workgroup = aws.athena.Workgroup("testWorkgroup", configuration={
    "resultConfiguration": {
        "encryption_configuration": {
            "encryptionOption": "SSE_KMS",
            "kms_key_arn": test_key.arn,
        },
    },
})
hoge_database = aws.athena.Database("hogeDatabase",
    bucket=hoge_bucket.id,
    name="users")
foo = aws.athena.NamedQuery("foo",
    database=hoge_database.name,
    query=hoge_database.name.apply(lambda name: f"SELECT * FROM {name} limit 10;"),
    workgroup=test_workgroup.id)


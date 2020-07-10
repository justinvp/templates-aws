import pulumi
import pulumi_aws as aws

strict = aws.iam.AccountPasswordPolicy("strict",
    allow_users_to_change_password=True,
    minimum_password_length=8,
    require_lowercase_characters=True,
    require_numbers=True,
    require_symbols=True,
    require_uppercase_characters=True)


import pulumi
import pulumi_aws as aws

example_user = aws.iam.User("exampleUser",
    force_destroy=True,
    path="/")
example_user_login_profile = aws.iam.UserLoginProfile("exampleUserLoginProfile",
    pgp_key="keybase:some_person_that_exists",
    user=example_user.name)
pulumi.export("password", example_user_login_profile.encrypted_password)


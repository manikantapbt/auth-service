# Auth-Service

Auth Service is a gRPC service responsible for handling user signups and logins. 
It ensures security by utilizing one-time password (OTP) verification sent to the
user's mobile. These OTPs are generated at runtime using Time-based One-Time Password
(TOTP), enhancing security as OTPs are not stored in the system. Additionally, the 
service tracks various user events such as login attempts and successes for analytics purposes.


The Application have the following APIs
 ### 1. SignupWithPhoneNumberRequest
    
This api stores the user information and sends the otp to user by publishing notification on rabbit mq to `otp-service`. 
The user needs to verify the phone number to mark user profile as verified.

input
```yaml
  int32 id = 1;
  string name = 2;
  string user_name = 3;
  string email = 4;
  bool isVerified = 5;
  int32 countryCode = 6;
  string PhoneNumber = 7;
```
output

```yaml
  bool isSuccess = 1;
  Error error = 2;
  int32 userId = 3;
```

### Features: 
1. Strong validation on user inputs
2. Stores user information to Postgres
3. Sends notification to otp-service to send otp to user's mobile number for verification
4. Logs SIGN_IN_REQUEST_OTP user event to user database.


### 2. VerifyPhoneNumber

This api validates the otp sent by notification service, the app uses time based One time password(totp) to validate the passwords. if the otp is correct then the user profile is marked as verified and
PHONE_VERIFIED user event is saved to db. if not the WRONG_OTP user event is logged to db

input

```yaml
  string requestId = 1;
  int32 otp = 2;
  int32 countryCode = 3;
  string phoneNumber = 4;
```
output

```yaml
  bool isSuccess = 1;
  Error error = 2;
```
### Features:
1. Strong validation on user inputs
2. Validates the sent OTP using totp
3. Logs PHONE_VERIFIED or WRONG_OTP user event to user database.

### 3. LoginWithPhoneNumber
Once the user is verified their phone number, the user can request login with phone number. The otp is sent to user mobile number for logging in. 


input

```yaml
  string requestId = 1;
  int32 countryCode = 2;
  string phoneNumber = 3;
```
output

```yaml
  bool isSuccess = 1;
  Error error = 2;
```
### Features:
1. Strong validation on user inputs
2. Sends notification to otp-service to send otp to user's mobile for login
3. Logs UNVERIFIED_LOGIN_ATTEMPT event to db if user tried to login without verified mobile number.
4. Logs LOGIN_REQUEST to db for verified profiles.

### 4. ValidatePhoneNumberLogin
This api validated the otp generated for user login.


input

```yaml
  string requestId = 1;
  int32 otp = 2;
  int32 countryCode = 3;
  string phoneNumber = 4;
```
output

```yaml
  bool isSuccess = 1;
  Error error = 2;
```
### Features:
1. Strong validation on user inputs
2. Validated the OTP generated by otp-service for user login using totp
3. Logs INCORRECT_OTP event to db if the sent otp is incorrect.
4. Logs LOGIN_SUCCESSFUL event to db if the sent otp is correct.


### 5. GetProfile
Retrieves the user information by user id

input

```yaml
  string requestId = 1;
  int32 userId = 2;
```

output

```yaml
  bool isSuccess = 1;
  Error error = 2;
  User user = 3;
```


### 6. GetProfileByPhoneNumber

Retrieves the user information by phone number.

input
```yaml
  string requestId = 1;
  int32 countryCode = 2;
  string phoneNumber = 3;
```
output
```shell
  bool isSuccess = 1;
  Error error = 2;
  User user = 3;
```

### Requirements

The app needs to run on atleast `go` version of `1.22`


```shell
 go version
```

### Install Dependencies
download dependencies into vendor folder
```shell
make mod
```

install dependencies
```shell
make install
```

### Running Tests

running tests with code coverage

```shell
make test
```

### Running the app
```shell
make mod
make run
```



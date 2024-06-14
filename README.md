# YoPAY

A Mini E-wallet integrated with a payment gateway.

## Features
- Get Information User
- Authentication (Login, Register)
- Charge Balance
- Get Balance
- Transfer Balance
- QR Code For Receiving Balance
- Withdraw Balance

## Tech Stack
- Golang
- Fiber
- Postgres
- Gorm
- Argon
- Midtrans

## Endpoints

1. **Authentication**:
    - **Register**: Endpoint for user registration. Path: `/v1/auth/register`
    - **Login**: Endpoint for user login. Path: `/v1/auth/login`

2. **User Information**:
    - **Get Current User**: Retrieves the current logged-in user's information. Path: `/v1/user/me`

3. **Balance Management**:
    - **Add Balance**: Allows users to add funds to their account. Path: `/v1/balance/add`
    - **Callback**: Endpoint for handling callbacks. Path: `/v1/balance/callback`
    - **Get Total Balance**: Retrieves the total balance of a user. Path: `/v1/balance/total`
    - **Transfer Balance**: Allows users to transfer funds from their account to another user's account. Path: `/v1/balance/transfer`
    - **Create Token Withdraw**: Creates a token for withdrawing funds. Path: `/v1/balance/withdraw`
    - **Confirm Withdraw**: Confirms the withdrawal of funds. Path: `/v1/balance/confirm-withdraw`


:label: I still need suggestions and feedback for the development of this application. If you want to contribute, please PM me on [Instagram](https://www.instagram.com/_imriann28/), or you can create an issue in this repository. Thank you.


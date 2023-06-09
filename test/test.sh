

# 用户注册

curl --location 'localhost:3000/auth/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "elon@musk.com",
    "passowrd": "1234567"

}'


# 用户登录

curl --location 'localhost:3000/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "elon@musk.com",
    "passowrd": "1234567"

}'
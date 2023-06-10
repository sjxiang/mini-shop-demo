

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


# 添加商品

curl --location 'localhost:3000/product/add' \
--header 'Authorization: bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTc5MDEzMzIsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MSwiRW1haWwiOiJlbG9uQG11c2suY29tIn0.0GyhMMZI43cF1J_E9HPB99U3Uvqm10cgMLaqWouJCaY' \
--header 'Content-Type: application/json' \
--data '{
    "name": "biyuntao",
    "stock": 20,
    "price": 5
}'


# 查询商品

curl --location --request GET 'localhost:3000/product/1' \
--header 'Authorization: bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTc5MDEzMzIsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MSwiRW1haWwiOiJlbG9uQG11c2suY29tIn0.0GyhMMZI43cF1J_E9HPB99U3Uvqm10cgMLaqWouJCaY' \



# 创建订单

curl --location 'localhost:3000/order/add' \
--header 'Authorization: bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTc5MDEzMzIsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MSwiRW1haWwiOiJlbG9uQG11c2suY29tIn0.0GyhMMZI43cF1J_E9HPB99U3Uvqm10cgMLaqWouJCaY' \
--header 'Content-Type: application/json' \
--data '{
    "product_id": 1,
    "quantity": 1
}'
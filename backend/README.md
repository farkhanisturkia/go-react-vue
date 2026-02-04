# Run Backend
1. go run .               → normal run
2. go run . reset         → reset users lalu run
3. go run . seed          → seed ulang lalu run
4. go run . reset-seed    → reset + seed lalu run

# Register
curl -X POST "http://localhost:3000/api/register"   -H "Content-Type: application/json"   -d '{
    "name": "Fika Ridaul Maulayya",
    "username": "maulayyacyber",
    "email": "fika@santrikoding.com",
    "password": "password"
}'

# Login
curl -X POST "http://localhost:3000/api/login"   -H "Content-Type: application/json"   -d '{
    "username": "maulayyacyber",
    "password": "password"
}'

# Get users
curl "http://localhost:3000/api/users?page=1&size=10" \
  -H "Authorization: Bearer ..."

# Create users
curl -X POST "http://localhost:3000/api/users" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ..." \
  -d '{
    "name": "Akhmad Lutfi",
    "username": "lutfi",
    "email": "lutfi@santrikoding.com",
    "password": "password"
}'

# Get user by ID
curl "http://localhost:3000/api/users/1" \
  -H "Authorization: Bearer ..."

# Update user by ID
curl -X PUT "http://localhost:3000/api/users/1" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ..." \
  -d '{
    "name": "Fika Ridaul Maulayya - Edit",
    "username": "maulayyacyber",
    "email": "fika@santrikoding.com",
    "password": "password"
}'

# Delete user by ID
curl -X DELETE "http://localhost:3000/api/users/1" \
  -H "Authorization: Bearer ..."
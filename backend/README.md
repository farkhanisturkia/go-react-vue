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
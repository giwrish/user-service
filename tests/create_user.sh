curl -X POST http://localhost:8080/api/user/ \
  -H "Content-Type: application/json" \
  -d '{
           "username": "test",
           "password": "securepassword"
         }' | jq

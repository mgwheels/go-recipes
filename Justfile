run:
  go run ./cmd/api

add-toast:
    curl -X POST http://localhost:8080/api/v1/recipes \
      -H "Content-Type: application/json" \
      -d @tmp/inputs/toast.json

add-gc:
    curl -X POST http://localhost:8080/api/v1/recipes \
      -H "Content-Type: application/json" \
      -d @tmp/inputs/grilled-cheese.json

list-recipes:
    curl http://localhost:8080/api/v1/recipes

update ID="1":
    curl -X PUT http://localhost:8080/api/v1/recipes/{{ID}} \
      -H "Content-Type: application/json" \
      -d @tmp/inputs/toast.json

delete ID="1":
    curl -X DELETE http://localhost:8080/api/v1/recipes/{{ID}}

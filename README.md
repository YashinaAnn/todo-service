# Todo service

## REST API

### Get todo list
Request:
```
GET /todos
```
Response:
```
[
    {
        "id": 4,
        "item": "Pass the exam",
        "completed": true
    },
    {
        "id": 5,
        "item": "Finish the book",
        "completed": false
    },
    {
        "id": 6,
        "item": "Go to gim",
        "completed": false
    }
]
```

### Get todo item by id
Request:
```
GET /todos/4
```
Response:
```
{
  "id": 4,
  "item": "Pass the exam",
  "completed": true
}
```

### Create todo item
Request:
```
POST /todos 

{
  "item": "Finish the book",
  "completed": false
}
```
Response:
```
{
  "id": 6,
  "item": "Finish the book",
  "completed": false
}
```


### Update task status
Request:
```
PATCH /todos/6
```
Response:
```
{
  "id": 6,
  "item": "Finish the book",
  "completed": true
}
```

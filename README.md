# TodoApp

### What is this project ?
 In this project our goal is learning fundemantals of golang programming language with using tools like postgres , kafka, redis,swagger linter,docker etc ... It has  two entities such as Todo and User. User can take 9 values but 6 of them given as default. Default values for User entitiy are ID, Todos(empty struct till you give values),IsEmailActive,CreatedAt,UpdatedAt and DeletedAt.Todos can take 7 values and 5 of them given as default. Default values for Todo entity are ID, UserID(can't be nil) ,CreatedAt UpdatedAt and DeletedAt.User and Todo are linked by one to many method.For example one user can have 3 todos but one todo can't have 2 users.
 
### HTTP Verbs For User

| HTTP METHOD |POST   | GET         | DELETE    | PATCH  |          
| ------------| ------|-------------|------------- | ------------- | 
|     |CREATE |READ    |DELETE       | UPDATE       | 
|      |/sign_up|/user/{user_id}|/user/delete/{user_id}|/user/update/{user_id}
||**-----------**|/users|**-----------**|/resetpassword
||**-----------**||**-----------**|/activation/{user_id}|

### HTTP Verbs For Todo

| HTTP METHOD |POST   | GET         | DELETE    | PATCH  |          
| ------------| ------|-------------|------------- | ------------- | 
|     |CREATE |READ    |DELETE       | UPDATE       | 
|      |/todo/create|/todo/:id|user/:userid/todo/delete/:todoid|user/:userid/todo/update/:todoid
||**-----------**|/todos/|**-----------**|**-----------**|

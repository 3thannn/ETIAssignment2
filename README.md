# ETIAssignment2
## API Endpoints:
| Microservice  | Port | Endpoint URL |
| ------------- | ---- | ------------ |
| Frontend  | 9040 | http://10.31.11.12:9040 |
| Comment API  | 9091 | http://10.31.11.12:9041 |
| Rating API  | 9092 | http://10.31.11.12:9042 |

## Documentation

### Comment Data Structure
| Attribute | Data Type |
| --------- | ---- |
| CommentID | uint |
| CreatorType | varchar(255) |
| CreatorID | varchar(9) |
| TargetType | varchar(255) |
| TargetID | varchar(9) |
| CommentData | varchar(500) |
| Anonymous | bool |
| DateTimePublished | datetime |

### Rating Data Structure
| Attribute | Data Type |
| --------- | ---- |
| RatingID | uint |
| CreatorType | varchar(255) |
| CreatorID | varchar(9) |
| TargetType | varchar(255) |
| TargetID | varchar(9) |
| RatingScore | varchar(500) |
| Anonymous | bool |
| DateTimePublished | datetime |

# Comment Microservice API Documentation
### [GET] /api/comment/{id}
Get comment by CommentID
```

Endpoint
http://10.31.11.12:9041/api/comment/{id}
Response
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```

### [PUT] /api/comment/{id}
Update comment by CommentID
```
Endpoint:
http://10.31.11.12:9041/api/comment/{id}
Response:
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```

### [POST] /api/comment
Create Comment
```
Endpoint:
http://10.31.11.12:9041/api/comment
Response:
Status code 201 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```


### [GET] /api/comment/received/{type}/{id}
Get received comments by user type and id
```
Endpoint:
http://10.31.11.12:9041/api/comment/received/{type}/{id}
Response:
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```

### [GET] /api/comment/received/anonymous/{type}/{id}
Get received anonymous comments by user type and id
```
Endpoint:
http://10.31.11.12:9041/api/comment/received/anonymous/{type}/{id}
Response:
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```

### [GET] /api/comment/posted/{type}/{id}
Get user's posted comments
```
Endpoint:
http://10.31.11.12:9041/api/comment/posted/{type}/{id}
Response:
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```

### [GET] /api/comment/student/{id}
### [GET] /api/comment/tutor/{id}
### [GET] /api/comment/class/{id}
### [GET] /api/comment/module/{id}
Get comments for all students, tutor, classes or modules.
```
Endpoint:
http://10.31.11.12:9041/api/comment/student/{id}<br />
http://10.31.11.12:9041/api/comment/tutor/{id}<br />
http://10.31.11.12:9041/api/comment/class/{id}<br />
http://10.31.11.12:9041/api/comment/module/{id}<br />
Response:
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```

## Sample of Comment jsonstring returned for GET request:
```
{
  "CommentID":1,
  "CreatorID":"1",
  "CreatorType":"Student",
  "TargetID":"2",
  "TargetType":"Student",
  "CommentData":"Great at working in teams.",
  "Anonymous":0,
  "DateTimePublished":"2022-02-04 04:14:06",
  "CreatorName":"Ethan",
  "TargetName":"Test"
 }
```

# Rating Microservice API Documentation

### [GET] /api/rating/{id}
Get Rating by RatingID
```
Endpoint:
http://10.31.11.12:9042/api/rating/{id}
Response:
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```


### [PUT] /api/rating/{id}
Update Rating 
```
Endpoint
http://10.31.11.12:9042/api/rating/{id}
Response
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```

### [POST] /api/rating/create
Create Rating
```
Endpoint:
http://10.31.11.12:9042/api/rating/create
Response:
Status code 201 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```


### [GET] /api/rating/received/{type}/{id}
Get received ratings by user type and id
```
Endpoint:
http://10.31.11.12:9042/api/rating/received/{type}/{id}
Response:
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```

### [GET] /api/rating/received/anonymous/{type}/{id}
Get received anonymous ratings by user type and id
```
Endpoint:
http://10.31.11.12:9042/api/rating/received/anonymous/{type}/{id}
Response:
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```

### [GET] /api/rating/posted/{type}/{id}
Get poasted ratings by user type and id
```
Endpoint:
http://10.31.11.12:9042/api/rating/posted/{type}/{id}
Response:
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```

### [GET] /api/rating/student/{id}
### [GET] /api/rating/tutor/{id}
### [GET] /api/rating/class/{id}
### [GET] /api/rating/module/{id}
Get ratings for all students, tutor, classes or modules.
```
Endpoint:
http://10.31.11.12:9042/api/rating/student/{id}<br />
http://10.31.11.12:9042/api/rating/tutor/{id}<br />
http://10.31.11.12:9042/api/rating/class/{id}<br />
http://10.31.11.12:9042/api/rating/module/{id}<br />
Response:
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```

## Sample of Rating jsonstring returned for GET request:
```
{"RatingID":1,
"CreatorID":"2",
"CreatorType":"Student",
"TargetID":"1",
"TargetType":"Student",
"RatingScore":3,
"Anonymous":1,
"DateTimePublished":"2022-02-04 04:14:06",
"CreatorName":"Kester",
"TargetName":"Ethan"}
```

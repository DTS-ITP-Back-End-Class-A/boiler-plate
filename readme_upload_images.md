

## Flow Upload Images

| Step        | Flow                 | 
|---------------|-----------------------|
| Pertama       | Hit Endpoint /upload           | 
| Kedua         | Validasi data images| 
| Ketiga        | show response with path images |

## Endpoint Review

| Name | Method | Endpoint | Description | 
|------|-------| ------- | ------- |
| upload images| POST | `/upload` | this endpoint will upload images and return path 

## Params

```
none
```

## Authorization

```
none
```

## Request Body

| key        | Value                 | 
|---------------|-----------------------|
| images_upload       | {{images}} | 

`Notes` : this request body will set in form-data

6. Response 

#### status 200
```
{
    "message": "success",
    "status": "ok",
    "data": {
        "uuid": "asjkd-asghdikj-ahsjdih",
        "images_url": "images/kajhsduihasd.jpg"
    }
}
```

#### status 400
```
{
    "message": "invalid",
    "status": "failed",
}
```

#### status 500
```
{
    "message": "Cannot connect DB",
    "status": "internal server error",
}
```
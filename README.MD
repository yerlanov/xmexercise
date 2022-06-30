# XM Exercise

config.yml is configuration file

    listen.port: is to configure application port

    db.uri: database connection

    ip_api_url: base api for ipapi.com

    jwt_secret: is token signature

Database migration provided in /internal/storage/migration folder

API examples
Create company

Endpoint requires bearer token, token can be generated on https://jwt.io
with signature. And method available on Cyprus

    POST /api/v1/company

    JSON body
    {
        "name":"name1",
        "code":"code2",
        "country":"country3",
        "website":"website4",
        "phone":"phone5"
    }

Update company


    PUT /api/v1/company/:id

    JSON body
    {
        "name":"name1",
        "code":"code2",
        "country":"country3",
        "website":"website4",
        "phone":"phone5"
    }

Delete company

Endpoint requires bearer token, token can be generated on https://jwt.io
with signature. And method available on Cyprus

    DELETE /api/v1/company/:id


List all

    GET /api/v1/company/list

    
List with filter

    GET /api/v1/company/filter?name=name1&phone=phone4
  

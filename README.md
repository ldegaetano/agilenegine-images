# Images

Images is a project that implements a transparent cache. It allows you to retrieve all the pictures saved in AgileEngine without making a lot of request to the AgileEngine API.

## Installation

Use docker to run the app.

```bash
docker build -t ae-images .
docker run -p 8080:8080 ae-images
```

## Usage

#### Request
```CURL
GET 'localhost:8080/images'
```
You can add a query param like this /images?page=2.

#### Response
- Status code: 200 OK
- Body:
````
{
    "pictures": [
        {
            "cropped_picture": "http://interview.agileengine.com/pictures/cropped/03 - WPIX_86.jpg",
            "id": "f6dbc60eb91e29426955"
        },
        {
            "cropped_picture": "http://interview.agileengine.com/pictures/cropped/04 - WPIX_86.jpg",
            "id": "08865e53ba703444389d"
        },
        {
            "cropped_picture": "http://interview.agileengine.com/pictures/cropped/04sc002.jpg",
            "id": "bc414f07c8b3b1880db5"
        },
        {
            "cropped_picture": "http://interview.agileengine.com/pictures/cropped/05wbx37rbaz40djvay2r9djygkgqb3mu_2.jpg",
            "id": "83a3f2fc22ea1f72f0be"
        },
        {
            "cropped_picture": "http://interview.agileengine.com/pictures/cropped/06 - WPIX_86.jpg",
            "id": "db8dbb43fef23062b234"
        },
        {
            "cropped_picture": "http://interview.agileengine.com/pictures/cropped/06sc001.jpg",
            "id": "b9eab2c7179eada11845"
        },
        {
            "cropped_picture": "http://interview.agileengine.com/pictures/cropped/07 - WPIX_86.jpg",
            "id": "7d8d8f6deebaa33832e5"
        },
        {
            "cropped_picture": "http://interview.agileengine.com/pictures/cropped/07sc011.jpg",
            "id": "92012604e03781b12451"
        },
        {
            "cropped_picture": "http://interview.agileengine.com/pictures/cropped/08 - WPIX_86.jpg",
            "id": "557e2419479173535992"
        },
        {
            "cropped_picture": "http://interview.agileengine.com/pictures/cropped/09sc049.jpg",
            "id": "6d78efcae2701d94c8aa"
        }
    ],
    "page": 2,
    "page_count": 26,
    "has_more": true
}
````

- Status code: 404 Not found

## Notes

- The project was developed in two hours. So there are many things that can be improved or added like error manage, logging, metrics, etc.
- There is only one handler implemented but all the service layer was implemented. So, you can see the service to check the algorithm used in the search function.

<h1>Forum API in Go</h1>

My first attempt at writing a REST API in Go, also using postgres 15.3 and docker. The server runs on localhost:8080 by default and the allowed origin is localhost:5173, both can be changed in main.go.

<h3>Database schema:</h3>

![image](https://github.com/LiamFarese/Forum-API-in-Go/assets/108936972/3273b52b-7f5a-46aa-8f4e-bacfc7465172)


<h3>Set-up:</h3>

have docker installed and running and Go 1.2 or newer installed

    Run the following commands:

        make pullpostgres
        make postgres
        make createdb
        make migrateup


<h3>Start:</h3>

To start server, make sure your docker container is running, then from project root:

    go run cmd/main.go

<h3>API Usage:</h3>

    User endpoints:

        /register - POST - takes following format json: {"username":"","password":"",role:""}
        /login - POST - takes following format json: {"username":"","password":""}

        /users - GET - returns all users on database
        /profile/{ID} - GET - ID is the id of the user, this returns the users profile which consists of their info, listings and posts

    Listing endpoints: 

        /listing - POST - creates a new listing, takes following json format: {"title":"","body":"","user_id":int}
        /listings - GET - returns all listings that are still available
        /listing/{ID} - GET - retuns a listing by its ID
        /listing/{ID} - PATCH - changes a listing from open to settled

    Post endpoints:

        /post - POST - creates a new post, takes following json format: {"title":"","body":"","user_id":int}
        /posts - GET - returns all posts
        /post/{ID} - GET - retuns a post by its ID

    Comment endpoints:

        /comment - POST - creates a new comment, takes following json: {"body":"","user_id":int,"post_id":int}
        /reply - POST - creates a reply to a comment, takes following json: {"body":"","user_id":int,"post_id":int, "parent_comment":int}
        /comments/{postId} - GET - returns all the comments for a post

<h3>Changes I'd like to make in the future:</h3>

* Include authentication with OAuth2





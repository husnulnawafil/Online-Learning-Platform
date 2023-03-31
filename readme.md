# ONLINE LEARNING PLATFORM ENDPOINTS

## How to

- ### run locally
1. Install Golang.
2. Install Golang Air [documentation here](https://github.com/cosmtrek/air).
3. Clone or fork repository.
4. Run `go mod tidy`.
5. Put `.env` file on root directory project.
6. Simply run `air init` to initiate the air.
7. Run `air` to run the app.
8. Download postman collection [postman collection](./docs/Online%20Learning%20Platform.postman_collection.json).
9. Send request to `localhost:<port>`.

<br>

- ### run on the server
1. Ensure that you have downloaded the [postman collection](./docs/Online%20Learning%20Platform.postman_collection.json)
2. Replace `http://127.0.0.1:3000` (localhost) to `http://34.143.167.48:3000` (hosting IP Address).
3. Pick request and click send on Postman or other http client.

<br>

## All you need to know
It is highly recommended to register then login first to get `USER` role account in order to use other endpoints. But please note that the `USER` role account only authorized to access a few endpoints so far, otherwise http will send status `403 Forbidden`. 
<br>
`ADMIN` role is needed to use all endpoint, please reach me on this [email](mailto:nawafilhusnul@gmail.com?subject=REQUEST%20ADMIN%20ACCOUNT%20-%20ONLINE%20PLATFORM%20LEARNING).


<br>

## Tech Stacks
1. Go (using [fiber](https://gofiber.io/)).
2. MongoDB (using VM local installation).
3. Google Cloud Platform (use [Compute Engine](https://cloud.google.com/compute)).
4. [Cloudinary](https://cloudinary.com/) for media upload.
5. JWT ([JSON Web Token](https://jwt.io/)) for aunthentication.
6. [Air](https://github.com/cosmtrek/air) - Live reload for Go apps.
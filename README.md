# Daily Word API

**Created a simple API that provides a daily word for to the user**

*To run the API:*
- clone the repository:
```bash
git clone https://github.com/plattnotpratt/daily-word-api.git
cd daily-word-api
```
-  Run the tidy command in go:
```bash
go mod tidy
```
- Run the server:
```bash
go run .
```
- Use Curl or the browser to navigate to an end point:
    - GET http://localhost:8080/daily-word - Returns a json response with a random word that only changes at the start of a new day. 12:00 AM
    - GET http://localhost:8080/random-word - Returns a json response with a single random word.
    - GET http://localhost:8080/random-word/4 - Returns a json response with a single random word with the length of 4
    - GET http://localhost:8080/random-words/3 - Returns an json response with an array of 3 words
    - GET http://localhost:8080/random-words/3/4 - Returns a json response with an array of 3 words that are all length 4


------------------------------------------------------------------------------------------

#### Introductory Endpoint to the API

<details>
 <summary><code>GET</code> <code><b>/</b></code> <code>(Get the introduction response to the API and ensure the API is running)</code></summary>

##### Parameters

> None

##### Responses

> | http code     | content-type                      | response                                                            |
> |---------------|-----------------------------------|---------------------------------------------------------------------|
> | `200`         | `text/plain;charset=UTF-8`        | `json String`                                                          |

##### Example cURL

> ```javascript
>  curl -X GET -H "Content-Type: application/json" http://localhost:8080/
> ```

</details>

------------------------------------------------------------------------------------------

#### Daily Word Endpoint

<details>
 <summary><code>GET</code> <code><b>/daily-word</b></code> <code>(Gets a word that changes only at the start of a new day based on Eastern Timezone)</code></summary>

##### Parameters

> None

##### Responses

> | http code     | content-type                      | response                                                            |
> |---------------|-----------------------------------|---------------------------------------------------------------------|
> | `200`         | `application/json`        | `json String`                                                          |
> | `400`       | `application/json`      | `json string`                                                        |

##### Example cURL

> ```javascript
>  curl -X GET -H "Content-Type: application/json" http://localhost:8080/daily-word
> ```

</details>

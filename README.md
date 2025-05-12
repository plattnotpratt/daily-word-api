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
 <summary><code>GET</code> <code><b>/daily-word</b></code> <code>(Gets a word that changes only at the start of a new day based on UTC-5:00)</code></summary>

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


------------------------------------------------------------------------------------------

#### Random Word Endpoints

<details>
 <summary><code>GET</code> <code><b>/random-word</b></code> <code>(Gets a random word every time the request is made.)</code></summary>

##### Parameters

> None

##### Responses

> | http code     | content-type                      | response                                                            |
> |---------------|-----------------------------------|---------------------------------------------------------------------|
> | `200`         | `application/json`        | `json String`                                                          |
> | `400`       | `application/json`      | `json string`                                                        |

##### Example cURL

> ```javascript
>  curl -X GET -H "Content-Type: application/json" http://localhost:8080/random-word
> ```

</details>

<details>
 <summary><code>GET</code> <code><b>/random-word/{length}</b></code> <code>(Gets a random words with the specified length)</code></summary>

##### Parameters

> | name              |  type     | data type      | description                         |
> |-------------------|-----------|----------------|-------------------------------------|
> | `length` |  required | int   | The number of characters in the random word.        |


##### Responses

> | http code     | content-type                      | response                                                            |
> |---------------|-----------------------------------|---------------------------------------------------------------------|
> | `200`         | `application/json`        | `json String`                                                          |
> | `400`       | `application/json`      | `json string`                                                        |

##### Example cURL

> ```javascript
>  curl -X GET -H "Content-Type: application/json" http://localhost:8080/random-word/4
> ```

</details>


<details>
 <summary><code>GET</code> <code><b>/random-words/{count}</b></code> <code>(Gets an array of random words based on the number specified )</code></summary>

##### Parameters

> | name              |  type     | data type      | description                         |
> |-------------------|-----------|----------------|-------------------------------------|
> | `count` |  required | int   | the number of words you would like returned in the array.(This does not avoid duplicates)        |


##### Responses

> | http code     | content-type                      | response                                                            |
> |---------------|-----------------------------------|---------------------------------------------------------------------|
> | `200`         | `application/json`        | `json String`                                                          |
> | `400`       | `application/json`      | `json string`                                                        |

##### Example cURL

> ```javascript
>  curl -X GET -H "Content-Type: application/json" http://localhost:8080/random-words/4
> ```

</details>

<details>
 <summary><code>GET</code> <code><b>/random-words/{count}/{length}</b></code> <code>(Gets an array of random words based on the number and filters them based on specified length)</code></summary>

##### Parameters

> | name              |  type     | data type      | description                         |
> |-------------------|-----------|----------------|-------------------------------------|
> | `count` |  required | int   | the number of words you would like returned in the array.(This does not avoid duplicates)        |
> | `length` |  required | int   | the number of characters in the array of words.(This does not avoid duplicates)        |



##### Responses

> | http code     | content-type                      | response                                                            |
> |---------------|-----------------------------------|---------------------------------------------------------------------|
> | `200`         | `application/json`        | `json String`                                                          |
> | `400`       | `application/json`      | `json string`                                                        |

##### Example cURL

> ```javascript
>  curl -X GET -H "Content-Type: application/json" http://localhost:8080/random-words/4/5
> ```

</details>

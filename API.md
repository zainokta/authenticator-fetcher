# Register

* **URL**

  /api/v1/user/register

* **Method:**
  
  `POST`

* **Data Params**

  ```json
    {
        "name": "string|required",
        "phone": "string|required",
        "role": "string|required|user,admin" 
    }
  ```

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{
        "data": "string",
        "status": "success"
    }`
 
* **Error Response:**

  * **Code:** 400 BAD REQUEST <br />
    **Content:** `{
        "data": "string",
        "status": "error"
    }`

  OR

  * **Code:** 500 INTERNAL SERVER ERROR <br />
    **Content:** `{
        "data": "string",
        "status": "error"
    }`

* **Notes:**

  After successful register, user will obtain a password from the response which will be used for login. 

# Login

* **URL**

  /api/v1/user/register

* **Method:**
  
  `POST`

* **Data Params**

  ```json
    {
        "phone": "string|required",
        "password": "string|required" 
    }
  ```

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{
        "data": "string",
        "status": "success"
    }`
 
* **Error Response:**

  * **Code:** 400 BAD REQUEST <br />
    **Content:** `{
        "data": "string",
        "status": "error"
    }`

  OR

  * **Code:** 500 INTERNAL SERVER ERROR <br />
    **Content:** `{
        "data": "string",
        "status": "error"
    }`

* **Notes:**

  After successful login, user will receive token from the response.

# Fetcher

* **URL**

  /api/v1/fetcher

* **Method:**
  
  `GET`

* **Request Header**
    ```
    Authorization: Bearer token
    ```

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{
        "data": "string",
    }`
 
* **Error Response:**

  * **Code:** 403 UNAUTHORIZED <br />
    **Content:** `{
        "data": "string",
    }`

  OR

  * **Code:** 500 INTERNAL SERVER ERROR <br />
    **Content:** `{
        "data": "string",
    }`

* **Notes:**

  On the first endpoint hit, it often takes time to request because this endpoint is make a request to external API, for the next request, redis will cache the external API data for 1 hour to make the request more responsive.


# Fetcher (Admin)

* **URL**

  /api/v1/admin/fetcher

* **Method:**
  
  `GET`

* **Request Header**
    ```
    Authorization: Bearer token
    ```

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{
        "data": "string"
    }`
 
* **Error Response:**

  * **Code:** 403 UNAUTHORIZED <br />
    **Content:** `{
        "data": "string",
    }`

  OR

  * **Code:** 500 INTERNAL SERVER ERROR <br />
    **Content:** `{
        "data": "string",
    }`

* **Notes:**

  This request is the same with normal request, the different is, this request is only for user with admin role.

# Verify JWT

  An endpoint to verify JWT Token.

* **URL**

  /api/v1/jwt/verify


* **Method:**
  
  `GET`

* **Request Header**
    ```
    Authorization: Bearer token
    ```

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** 
    ```json
    {
        "data": {
            "CreatedAt": "string",
            "Name": "string",
            "Phone": "string",
            "Role": "string",
            "exp": "number",
            "iss": "string"
        },
        "status": "success"
    }
    ```
 
* **Error Response:**

  * **Code:** 400 BAD REQUEST <br />
    **Content:** `{
        "data": "string",
        "status": "error"
    }`

  OR

  * **Code:** 500 INTERNAL SERVER ERROR <br />
    **Content:** `{
        "data": "string",
        "status": "error"
    }`

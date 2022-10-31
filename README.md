# Imperial fleet inventory

## How to run project.

- run `make env`, to copy `.env.example` to `.env`
- run `make start-db` to start database
- run `make start-services` to start services

## Project endpoints

- `GET`:`http://http:8080/spaceships` - get all spaceships
    - Query params:
        - `name`
        - `class`
        - `status`
    - Example: `http://http:8080/spaceships?name=Devastator&status=operational&class=Star%20Destroyer`
    - Output
    ```json
    {
      "data": [
        {
          "id": 2,
          "name": "Devastator",
          "status": "operational"
        }
      ]
    }
    ``` 

- `GET`:`http://http:8080/spaceships/{id}` - get spaceship by id
    - Example: `http://http:8080/spaceships/2`
    - Output
    ```json
   {
      "id": 2,
      "name": "Devastator",
      "class": "Star Destroyer",
      "crew": 35000,
      "image": "https:\\\\url.to.image",
      "value": 1999.99,
      "status": "operational",
      "armament": [
        {
          "title": "Turbo Laser",
          "message_count": "60"
        }
      ]
    }
    ```

- `POST`:`http://http:8080/spaceships` - create spaceship
    - Request body
    ```json
    {
      "name": "Devastator2",
      "class": "Star Destroyer",
      "crew": 35000,
      "image": "https:\\\\url.to.image",
      "value": 1999.99,
      "status": "2operational",
      "armament": [
        {
          "title": "Turbo Laser",
          "message_count": "60"
        }
      ]
    }
    ```
    - Output
    ```json
    {
      "success": true
    }
    ```

- `PUT`:`http://http:8080/spaceships/{id}` - update spaceship
    - Request body
    ```json
    {
      "name": "Devastator",
      "class": "Star Destrowyer",
      "crew": 35000,
      "image": "https:\\\\url.to.image",
      "value": 1999.99,
      "status": "operational",
      "armament": [
        {
          "title": "Turbo Las22er",
          "message_count": "60"
        },
        {
          "title": "Turbo Laser2",
          "message_count": "60"
        }
      ]
    }
    ```
    - Output
    ```json
    {
      "success": true
    }
    ```

- `DELETE`:`http://http:8080/spaceships/{id}` - delete spaceship
  - Example: `http://http:8080/spaceships/2`
  - Output
  ```json
  {
    "success": true
  }
  ```

My attempt to create a REST API on Go with using Gin framework

# API Route Documentation

This document provides an overview of all API endpoints available in the application, along with HTTP methods, endpoint paths, accessability, and descriptions of their functionalities.

---

## User Routes

- **`POST /register`**
  - **Access**: Anyone
  - **Description**: Registers a new user.

- **`POST /login`**
  - **Access**: Anyone
  - **Description**: Authenticates a user and provides a token.

- **`PUT /user/:id`**
  - **Access**: Authorized Users
  - **Description**: Updates user information by user ID.

---

## Anime Routes

- **`GET /anime`**
  - **Access**: Anyone
  - **Description**: Retrieves a list of all anime.

- **`GET /anime/id/:id`**
  - **Access**: Anyone
  - **Description**: Retrieves anime details by ID.

- **`GET /anime/title/:title`**
  - **Access**: Anyone
  - **Description**: Retrieves anime details by title.

- **`GET /anime/rating/:id`**
  - **Access**: Anyone
  - **Description**: Retrieves rating of an anime by ID.

- **`POST /anime`**
  - **Access**: Administration
  - **Description**: Adds a new anime to the database.

- **`POST /anime/rating/:id`**
  - **Access**: Authorized Users
  - **Description**: Rates an anime by ID.

- **`PUT /anime/:id`**
  - **Access**: Administration
  - **Description**: Updates an anime entry by ID.

- **`DELETE /anime/:id`**
  - **Access**: Administration
  - **Description**: Deletes an anime entry by ID.

---

## Character Routes

- **`GET /characters`**
  - **Access**: Anyone
  - **Description**: Retrieves all characters.

- **`GET /characters/:id`**
  - **Access**: Anyone
  - **Description**: Retrieves a character by ID.

- **`GET /characters/anime/:id`**
  - **Access**: Anyone
  - **Description**: Retrieves characters related to a specific anime ID.

- **`POST /characters`**
  - **Access**: Administration
  - **Description**: Adds a new character.

- **`PUT /characters/:id`**
  - **Access**: Administration
  - **Description**: Updates a character by ID.

- **`DELETE /characters/:id`**
  - **Access**: Administration
  - **Description**: Deletes a character by ID.

---

## List Routes

- **`GET /list/anime`**
  - **Access**: Anyone
  - **Description**: Retrieves all anime lists.

- **`GET /list/anime/details`**
  - **Access**: Anyone
  - **Description**: Retrieves all anime lists with their content.

- **`GET /list/characters`**
  - **Access**: Anyone
  - **Description**: Retrieves all character lists.

- **`GET /list/characters/details`**
  - **Access**: Anyone
  - **Description**: Retrieves all character lists with their content.

- **`GET /list/anime/:id`**
  - **Access**: Anyone
  - **Description**: Retrieves an anime list by ID.

- **`GET /list/characters/:id`**
  - **Access**: Anyone
  - **Description**: Retrieves a character list by ID.

- **`GET /list/anime/anime/:id`**
  - **Access**: Anyone
  - **Description**: Retrieves anime lists by a specific anime ID.

- **`GET /list/characters/character/:id`**
  - **Access**: Anyone
  - **Description**: Retrieves character lists by a specific character ID.

- **`POST /list/anime`**
  - **Access**: Authorized users
  - **Description**: Creates a new anime list.

- **`POST /list/characters`**
  - **Access**: Authorized users
  - **Description**: Creates a new character list.

- **`POST /list/anime/add/:id`**
  - **Access**: Authorized users
  - **Description**: Adds an anime to a list by list ID.

- **`POST /list/characters/add/:id`**
  - **Access**: Authorized users
  - **Description**: Adds a character to a list by list ID.

- **`PUT /list/anime/edit/:id`**
  - **Access**: Authorized users
  - **Description**: Edits an anime list by ID.

- **`PUT /list/characters/edit/:id`**
  - **Access**: Authorized users
  - **Description**: Edits a character list by ID.

- **`POST /list/rating/:id`**
  - **Access**: Authorized users
  - **Description**: Updates the rating for a list by list ID.

- **`DELETE /list/anime/delete/:id`**
  - **Access**: Authorized users
  - **Description**: Deletes anime Lists and All Scores and Comments by List ID.

- **`DELETE /list/characters/delete/:id`**
  - **Access**: Authorized users
  - **Description**: Deletes characters Lists and All Scores and Comments by List ID.
---

## Comment Routes

- **`GET /comment`**
  - **Access**: Anyone
  - **Description**: Retrieves all comments.

- **`GET /comment/type/:type`**
  - **Access**: Anyone
  - **Description**: Retrieves comments by type.

- **`GET /comment/id/:id`**
  - **Access**: Anyone
  - **Description**: Retrieves a comment by ID.

- **`GET /comment/:type/:id`**
  - **Access**: Anyone
  - **Description**: Retrieves comments associated with a specific type and ID.

- **`GET /comment/user/:id`**
  - **Access**: Anyone
  - **Description**: Retrieves comments associated with a specific user.

- **`POST /comment`**
  - **Access**: Authorized users
  - **Description**: Posts a new comment.

- **`PATCH /comment/id/:id`**
  - **Access**: Authorized users
  - **Description**: Updates a comment by ID.

- **`POST /comment/rating/:id`**
  - **Access**: Authorized users
  - **Description**: Updates the rating of a comment by ID.

- **`DELETE /comment/id/:id`**
  - **Access**: Authorized users
  - **Description**: Deletes a comment by ID.

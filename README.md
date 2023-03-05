# Recipe 2 Grocery List - Backend

Link to [Frontend](client/README.md)

Simple app to store recipes and create grocery list from the ingredients of selected recipes.

![GO](https://img.shields.io/badge/go-1.20.1-blue.svg)

## Config

Uses `dotenv` to read env variables
- `SLACK_CHANNEL` - Slack channel ID to publish the grocery list
- `SLACK_BOT_TOKEN` - Slack bot Token
- `MONGODB_CLIENT` - MongoDB client path
- `MONGODB_DB` - MongoDB database name

```dotenv
SLACK_CHANNEL=""
SLACK_BOT_TOKEN=""
MONGODB_CLIENT=""
MONGODB_DB=""
```

## Endpoints
- POST `/recipe` - Add new recipe with the following params
  - `title` {`String`} - title of the recipe
  - `description` {`String`} - description of the recipe
  - `ingredients` {`Ingredient[]`} - list of ingredients
    - `name` {`String`} - name of ingredient
    - `amount` {`Float32`} - amount of ingredient
    - `unit` {`String`} - unit of ingredient amount
- GET `/recipes/${pageNum}` - Get paginated recipes
- GET `/recipes/${recipeId}` - Get recipe based on given id
- DELETE `/recipes/${recipeId}` - Remove recipe based on given id
- POST `/recipe/send` - Send ingredients of chosen recipes (backend could send a Slack notification etc)
  - `list` {`String[]`} - Grocery list
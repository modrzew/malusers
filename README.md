# malusers

Go scraper created to answer the ancient question: *who has the biggest power level?*

Uses [PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery) for parsing DOM and [jinzhu/gorm](https://github.com/jinzhu/gorm) for database manipulation.

## Installation

1. Clone:

    ```
    git clone git@github.com:modrzew/malusers
    cd malusers
    ```

2. Build Go code:

    ```
    go get ./...
    go build ./...
    ```

3. Build JS code:

    ```
    cd web
    npm install
    npm run build
    ```

4. Copy `config.json.example` to `config.json` and fill it with connection values to your Postgres database.

## Running

There are three CLI commands that can be run: `malapi`, `maldata` and `malscraper`, as well as web page.

All of them will read `config.json`, and use the same database.

### `malapi`

API for the web application.

### `maldata`

Used to process data in the database: create ranking table, as well as global statistics.

### `malscraper`

Fetches users' data from the Internet.

### Web page

#### Production

After running `npm run build`, production bundle will be placed in `build` folder.

#### Development

Run `npm run start`.

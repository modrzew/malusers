# malusers

Simple Go scraper created to answer the ancient question: *who has the biggest power level?*

Uses [asciimoo/colly](https://github.com/asciimoo/colly) for scraping and [jinzhu/gorm](https://github.com/jinzhu/gorm) for database manipulation.

## Usage

1. Clone and build:

    ```
    git clone git@github.com:modrzew/malusers
    cd malusers
    go get
    go build
    ```

2. Copy `config.json.example` to `config.json` and fill it with connection values to your Postgres database.
3. Run `./malusers`.

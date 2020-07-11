# ShortURL
ShortURL is crossplatform HTTP Server written in Go, offering static URL storage. It offers public and private URLs - public are listed on customisable HTML page.

## Usage
Running server without options leads to random port and `urls.csv` file for URL storage beeing chosen. Server provides options to customize those defaults:
- `-port [integer]` to set port where server listens
- `-csv-file [path]` to set source of url's and their short identifiers.

### Customisation
Edit `html.go` file to modify public html code.

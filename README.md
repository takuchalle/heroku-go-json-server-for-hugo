# OGP Parse Server for Hugo

## Usage on Heroku

```
$ go get github.com/takuchalle/heroku-go-json-server-for-hugo
$ cd $GOPATH/src/github.com/takuchalle/heroku-go-json-server-for-hugo
$ heroku login
$ heroku create
$ git push heroku master
```
And access `[your-own-herokuapp].herokuapp.com/api/ogp?url=https://hogehoge.com` with a browser.

## Usage local

```
$ go get github.com/takuchalle/heroku-go-json-server-for-hugo
$ cd $GOPATH/src/github.com/takuchalle/heroku-go-json-server-for-hugo
$ export PORT=1333
$ go run main.go
```

And access `localhost:1333/api/ogp?url=https://hogehoge.com` with a browser.

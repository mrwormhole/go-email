# go-email
Email microservice that works with sendgrid api for email sending and it has a view page that shows previous emails between people 🐳

##### To use this, run the command down below. Default selected database is sqlite3 for now. Make sure to change sendgrid api key and salt key in dockerfile before you do this
```
$ git clone github.com/MrWormHole/go-email.git
$ docker-compose up -d
```

##### POST /api/sendEmail [jwtAuth required, use the token(lasts 1 hour) when printed on console after application started]
##### request:
```
{
 "from": "test@test.com",
 "to" : "talhaaltinel@hotmail.com",
 "subject": "DD",
 "plaintext": "such a nice cat",
 "htmlcontent": "<h1>I LOVE YOUR CAT</h1>"
}
```
##### response(if successful):
```
{
 "from": "test@test.com",
 "to" : "talhaaltinel@hotmail.com",
 "subject": "DD",
 "plaintext": "such a nice cat",
 "htmlcontent": "<h1>I LOVE YOUR CAT</h1>"
}

```

##### GET /api/emails/:id [jwtAuth required, use the token(lasts 1 hour) when printed on console after application started]
```
  This will get a email from email records.
```
##### GET /api/emails [jwtAuth required, use the token(lasts 1 hour) when printed on console after application started]
```
  This will get all emails from email records.
```
##### DELETE /api/emails/:id [jwtAuth required, use the token(lasts 1 hour) when printed on console after application started]
```
  This will delete the email from email records.
```
##### GET /views/emails [basicAuth required, default name: jack, default password:1234]
```
  This will show you all email records in a view.
```
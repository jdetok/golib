# maild documentation: mime with standard smtp pkg
## sending with gmail
- gmail account has to have 2FA enabled
- after enabling, go to the `app passwords` section in acct settings & generate a new app password
# how to use 
## create .env file in root directory following this template: 
```
GMAIL_HOST=smtp.gmail.com
GMAIL_PORT=587
GMAIL_SNDR=example@gmail.com
GMAIL_PASS=zzzzxxxxccccvvvv
```
# sending options
## SendBasicEmail() 
- sends email without any attachements, only subject & text
## SendMIMEEmail(filename)
- sends an email with an attached file. pass the file name (has to exist), it will read & encode the file as a MIME compatible string then attach manually with MIME

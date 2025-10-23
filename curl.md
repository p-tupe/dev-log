---
modified: Thu Oct 23 12:48:34 EDT 2025
---
## curl

- Send an email using smtp

  ```sh
  curl --ssl-reqd \
    --url smtp://smtp.mail.com \
    --mail-from sender@mail.com \
    --mail-rcpt recipient@mail.com \
    --user "sender@mail.com:password" \
    --upload-file mail.txt
  ```

  where mail.txt:

  ```txt
  From: Sender Name <sender@mail.com>
  To: Recipient Name <recipient@mail.com>
  Subject: an example.com example email
  Date: Mon, 5 Nov 1994 12:10:00

  Dear Recipient,
  This is an example email.
  ```

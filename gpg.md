## gpg

- Encrypt file for self

  ```sh
  gpg --encrypt ["name" | --recipient <self@mail.com>] <file>
  ```

  This will create an encryped `<file>.gpg` file in same directory. Opening it normally will show gibberish.

- Decrypt a file

  ```sh
  gpg --decrypt <file>.gpg
  ```

  This will ask for your gpg passphrase and output the original `<file>`.

- Export public key to server (so can encrypt files there)

  ```sh
  gpg --armor --export <self@mail.com> > pub.key
  ```

  - Copy `pub.key` to server (via scp or copy-paste), then

  ```sh
  gpg --import pub.key
  ```

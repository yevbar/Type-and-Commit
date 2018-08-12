# Type and Commit
Made at Stupid Shit Hackathon

Credit to [MarinX's keylogger in golang](https://github.com/MarinX/keylogger)

To use with `vim`, go ahead and run the following

```bash
$ go get github.com/MarinX/keylogger
$ go build editor.go
$ sudo ./editor filename
```

_**Root permissions are needed to listen to key events**_

To simply use the keylogger program, go ahead and follow these instructions instead

```bash
$ go get github.com/MarinX/keylogger
$ go build logncommit.go
$ sudo ./logncommit
```

When done, <kbd>Ctrl + C</kbd> and push your new commits!

```bash
$ sudo git push origin master
```

_**Root permissions are needed since we commit under root**_

## Setting up permanent log-in in Linux

If you're like me and clone via https, then you'll need to save your credentials somewhere in order to use this.

Your options are to do one of the following (you don't need to do both!)

* [Cache your credentials for limited time](https://github.com/MarinX/keylogger)
* [Create a credentials file in a directory](https://github.com/MarinX/keylogger)

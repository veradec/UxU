
# UxU

UxU is a file hosting service, an alternative to [0x0.st](https://0x0.st) (which i am pretty sure isnt working) is written in Go.  
## Run Locally

Clone the project

```bash
  git clone https://git.veradec.xyz/UxU.git
```

Go to the project directory

```bash
  cd UxU
```

Install dependencies

```bash
  go get git.veradec.xyz/UxU
  go install
```

Start the server

```bash
  go run . --domain=https://yourdomain.xyz
```


## Demo

Insert gif or link to demo


## FAQ

#### There is no Upload Button, How do I actually Upload?

Using the curl command   
```curl -X POST https://uxu.veradec.xyz -F "file=@file.png"```

#### How do I Fetch the Files then?

the reponse will be a link with your Filename 

https://uxu.veradec.xyz/{random_string}/file.png

and you can use the link to view the file

## Acknowledgements

 - [0x0.st](https://0x0.st)
 - [Gin Documentation](https://gin-gonic.com/en/docs/)


## Feedback

If you have any feedback, please reach out to us at veradec@duck.com


## Contributing

Contributions are always welcome!

Mail me the git diffs - verade@duck.com

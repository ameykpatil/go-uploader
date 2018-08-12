# Go-Uploader
The service is helpful for uploading media & other assets to s3.   
Before diving into **go-uploader** first you will need to setup Go & **go-uploader** environment properly.

### Flow
Flow of the **go-uploader** is quite simple as follows -  
`router --> controller --> model`  
To store models hosted `MongoDB` is used.  
Controller also deals with `s3` calls for getting signed urls.  
Apart from these there are utilities & constant files in the project.  

  

### Install Go
You can skip this step if you have Go installed already.
```sh
brew update & brew install go           # install Go 
mkdir $HOME/gospace                     # workspace for Go development
export GOPATH=$HOME/gospace             # export $GOPATH variable
mkdir $GOPATH/src                       # contains Go source files
mkdir $GOPATH/pkg                       # contains package objects
mkdir $GOPATH/bin                       # contains executables
export PATH=$PATH:$GOPATH/bin           # export $GOPATH/bin
```
### Clone Project
It is better to follow standard directory structure for Go projects. This is important to avoid package imports complexities. It also keeps the Go workspace clean.
```sh
mkdir -p $GOPATH/src/github.com/ameykpatil                # create directory inside $GOPATH
cd $GOPATH/src/github.com/ameykpatil                      # go to the directory
git clone https://github.com/ameykpatil/go-uploader       # clone go-uploader inside the directory
```
### Test Project
There is `check.sh` file written which takes care of all the testing aspects from linting to tests & much more.  
There are integration as well as unit tests written.
```sh
# run tests
AWS_ACCESS_KEY_ID=<aws-access-key-id> AWS_SECRET_ACCESS_KEY=<aws-secret-access-key> AWS_REGION=us-east-1 AWS_BUCKET=ts-engineering-test MONGO_URL=<mongo-url> MONGO_DB_NAME=dblabs MONGO_COLLECTION=assets sh check.sh

# If you get any error regarding permissions, give executable permissions to check.sh & then run above command.
chmod +x check.sh

# check.sh checks for 4 aspects - formatting, correctness, linting & tests
# it might look slow as it must be connecting to remote database & services
```
### Run Project
```sh
# build go project (This creates binary & put it in $GOPATH/bin)
go install

# run go-uploader directly (as $GOPATH/bin is already added in $PATH)
AWS_ACCESS_KEY_ID=<aws-access-key-id> AWS_SECRET_ACCESS_KEY=<aws-secret-access-key> AWS_REGION=us-east-1 AWS_BUCKET=ts-engineering-test MONGO_URL=<mongo-url> MONGO_DB_NAME=dblabs MONGO_COLLECTION=assets go-uploader

# Check if server is running by opening following url, you should get response as "pong"
http://localhost:8489/ping
```
### Call APIs
You can use any web client to call APIs (such as `Postman`).  
But for keeping it simple here are the curl commands for the APIs.
```sh
# call post api to get upload_url along with asset-id
curl -X POST 'http://localhost:8489/asset'

# some shells might show unicode characters for some of the special characters in pre-signed s3 url, if that is the case copy upload_url & paste it in this code snippet & run, the output will give you proper utf-8 equivalent  
https://play.golang.org/p/-f5LaNmsCFg

# upload the asset to s3 using pre-signed url
curl -v --upload-file <filepath-to-upload> "<signed-upload-url>"

# update status of the asset using put api
curl -X PUT 'http://localhost:8489/asset/<asset-id>' -H 'Content-Type: application/json' -d '{"status": "uploaded"}'

# get the download_url for the asset given asset-id
curl -X GET 'http://localhost:8489/asset/<asset-id>?timeout=3600'

# download asset using returned download_url 
# make sure the extension is same that of the uploaded asset (.jpg) Also do not miss the quotes.  
curl "<signed-download-url>" --output sample.jpg
```

### Built With
* [Dep](https://github.com/golang/dep) - Dependency Management
* [Gin](https://github.com/gin-gonic/gin) - Web framework
* [Goblin](https://github.com/franela/goblin) - Testing
* [AWS-SDK](github.com/aws/aws-sdk-go) - S3 interaction
* [Mgo](https://gopkg.in/mgo.v2) - MongoDB interaction

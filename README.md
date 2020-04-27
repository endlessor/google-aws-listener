### Google RTB (real time bidding)
Project is written using `go1.13.3` and [Gin](https://github.com/gin-gonic/gin) framework. 

## About
This app streams Real-Time Bidding data from google to AWS S3 bucket.

## Configuration
All configs are in `config/config.json` file.  
In order to be able to proceed with file upload to S3 storage you need to configure aws credentials locally.  
Please run `aws configure`. You may also need to install `awscli`.

## Run application

There is Makefile which has predefined commands.  
Please run `make run` to start app locally.

## Performance testing

You can run performance tests using `ab` util

`ab -n 2000 -c 50 -p request.txt -T POST http://localhost:9000/api/rtb`

Where `n` stands for number of requests to perform and `c` stands for number of multiple requests to make at a time.  
Parameter `p`  is for file with json request body.  
Here is an example of `request.txt` file content:
```
{
	"number": 100,
	"string": "hello world",
	"float": 5.55
}
```

## Production deployment
In order to run application in production mode please set up these environment variables:  
`export GIN_MODE=release`  
`export WORKING_DIR="/var/www/go/"` full path to working directory.  
Don't forget a trailing slash.

## Bidder configuration
Go to this [page](https://developers.google.com/authorized-buyers/apis/v1.4/accounts#bidderLocation.bidProtocol) if you want to add more bidder listeners  
or get more up to date information on implementations options.

#### Add new bidder listener (example): 
Add access token to header in order to be able to make this request.  
```
PATCH https://www.googleapis.com/adexchangebuyer/v1.4/accounts/{ACCOUNT_ID}
```
With request body:
```json
{
	   "bidderLocation": [
        {
            "url": "https://www.google-rtb-test.com/api/rtb",
            "maximumQps": 250,
            "region": "US_EAST",
            "bidProtocol": "PROTOCOL_OPENRTB_2_3"
        },
        {
            "url": "https://singapore.google-rtb-test.com/api/rtb",
            "maximumQps": 250,
            "region": "ASIA",
            "bidProtocol": "PROTOCOL_OPENRTB_2_3"
        }
    ]
}
```
#### List all bidder listeners:
```
GET https://www.googleapis.com/adexchangebuyer/v1.4/accounts?access_token={ACCESS_TOKEN}
```

## Logger
Logs are stored by default in `log/app.log` file.  
This can be changed in config file.

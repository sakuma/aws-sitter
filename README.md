# AWS (EC2, RDS) start/stop control on AWS Lambda


## Using


### setup deploy tool

using [Serverless Framework](https://serverless.com/)

https://serverless.com/framework/docs/getting-started/

### AWS creadential
TODO:

### deploy

`$ make deploy`

※ see [Makefile](https://github.com/sakuma/aws-sitter/blob/master/Makefile)


### Instance setting (EC2, RDS)

maintanance `tag`

* API_CONTROLLABLE
  * `true` or `false`
* API_AUTO_OPERATION_MODE
  * `auto` : Activate during the time period specified in _API_RUN_SCHEDULE_.
  * `start` : Start only behavior. The start time is the start time specified in _API_RUN_SCHEDULE_.
  * `stop` : Stop only behavior. The stop time is the stop time specified in _API_RUN_SCHEDULE_.
* API_RUN_SCHEDULE
  * [time range(Hour)]:[weekday num (Sun:0, Mon:1)]
    * ex) 10-20  -> 10:00am 〜 20:59pm on every day
    * ex) 8-19:1-5  -> 8:00am 〜 19:59pm on Mon 〜 Fri
    * ex) 7-8:0,1,5  -> 7:00am 〜 8:59am on Sun, Mon, and Fri

### Note
Run every [30 minutes](https://github.com/sakuma/aws-sitter/blob/master/serverless.yml#L42).


## Development

### setup

* [Serverless Framework](https://serverless.com/)
* [Golang](https://golang.org/doc/install)
  * GO111MODULE=on
  * Version: see [go.mod](https://github.com/sakuma/aws-sitter/blob/master/go.mod)

### build

make build

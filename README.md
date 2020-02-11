# AWS (EC2, RDS) start/stop control on AWS Lambda


## setup

### DevOps: lambda

using [Serverless Framework](https://serverless.com/)

#### setup deploy tool

https://serverless.com/framework/docs/getting-started/


```
$ npm install -g serverless
```

Or, update the serverless cli from a previous version

```
$ npm update -g serverless
```

### setting file (instance list)

`$ cp instances.yml{.sample,}`

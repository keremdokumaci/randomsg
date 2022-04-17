# Randomsg

## A CLI tool to generate random messages and publish to cloud services like (SQS,SNS,PUB/SUB and etc.)

**Randomsg** is a basic CLI tool. Simply, adding message format, rules and required credentials for the service (like SNS) are enough to use it.

## Installation

randomsg requires go with version 1.18+.

Install go and run:

```sh
go install randomsg
```

After installation of randomsg, ready to use. Just run :

```sh
randomsg **args**
```

## Args

_This table will be uptaded for new features._

| Parameter | Need to know                       | Description                                            |
| --------- | ---------------------------------- | ------------------------------------------------------ |
| --service | sqs and sns are supported now.     | service which the message will be sent.                |
| --file    | .json files are supported now.     | file which has message format and service credentials. |
| --count   | should be greater than 0.          | count of messages to publish                           |
| --delay   | should be greater than or equaş 0. | delay before each message publish.                     |

## Sample Message File

You can find the sample in **master** branch as _sample.json_. All of the supported rules and types will be in that json.

#### Types

- string
- int
- time
- float
- char

#### Rules

- _min_ & _max_ for numeric types
- _format_ for string type.
  - available values for _format_ : **uuid**,**startsWith**,**endsWith**

## Contribution

Feel free to contribute !

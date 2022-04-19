# Randomsg

A CLI tool to generate random messages and publish to cloud services like (SQS,SNS,PUB/SUB and etc.).

**TODO**

Generation of nested objects is not supported yet. However, it is gonna be added as soon as possible !

## Installation

randomsg requires go with version 1.18+.

Install go and run:

```sh
go install "github.com/keremdokumaci/randomsg"@latest
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
- All numeric fields (int,float ..)
- time
- bool

#### Rules

- _min_ & _max_ for numeric types
- _startsWith_ & _endsWith_ for string type.
- _format_ for string type.
  - available values for _format_ : **uuid**

## Contribution

Feel free to contribute !

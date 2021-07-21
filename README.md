# Birdview

Birdview was built to facilitate management of large cloud estates by providing a one pane of glass view of the service and resources status.

# How to

Run

    make run

Build

    make test build


# Configuration

## AWS IAM

Create an IAM role with the following permissions 

```json
{
    "Version": "2012-10-17",
    "Statement": [{
        "Sid": "v1-20210712",
        "Effect": "Allow",
        "Action": [
          "health:Describe*",
          "config:List*",
          "organizations:List*",
          "organizations:Describe*",
          "config:BatchGet*"
        ],
        "Resource": "*"
    }]
}
```




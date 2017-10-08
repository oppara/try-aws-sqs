#!/bin/bash
set -eu

. $(cd $(dirname $0) && pwd)/.env

function delete_message() {
  aws sqs delete-message \
    --queue-url "${SQS_QUEUE_URL}" \
    --receipt-handle ${1}

  echo "delete MessageId: ${2}"
}

function execute() {
  while read -r line
  do
    message_id=$(echo "$line" | jq -r '.MessageId')
    body=$(echo "$line" | jq -r '.Body')
    handle=$(echo "$line" | jq -r '.ReceiptHandle')

    echo "MessageId: ${message_id}"
    echo "Body: ${body}"

    delete_message $handle $message_id
  done
}


aws sqs receive-message \
  --queue-url "${SQS_QUEUE_URL}" \
  --max-number-of-messages 10  | jq -rc '.Messages[]' | execute


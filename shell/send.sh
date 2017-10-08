#!/bin/bash
set -eu

. $(cd $(dirname $0) && pwd)/.env

for i in {0..4};
do
  id=$(uuidgen)
  body="body ${i}"
  message='{"id":"'$id'","body":"'$body'"}'

  aws sqs send-message \
    --queue-url "${SQS_QUEUE_URL}" \
    --message-body "${message}" \
    --message-group-id "shell"
done




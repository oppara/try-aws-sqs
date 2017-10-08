<?php
require_once __DIR__ . '/bootstrap.php';

try {
    $client = createClient();

    $result = $client->receiveMessage([
        'QueueUrl' => QUEUE_URL,
        'MaxNumberOfMessages' => 10,
    ]);

    $messages = $result->search('Messages[]') ?? [];
    foreach ($messages as $message) {
        $handle = $message['ReceiptHandle'];
        $body = $message['Body'];
        echo $body . PHP_EOL;
        $client->deleteMessage([
            'QueueUrl' => QUEUE_URL,
            'ReceiptHandle' => $handle,
        ]);
    }

} catch (\Aws\Sqs\Exception\SqsException $e) {
    var_export($e->getMessage());
}



<?php
require_once __DIR__ . '/bootstrap.php';

try {
    $client = createClient();

    for ($i = 0; $i < 5; $i++) {
        $client->sendMessage([
            'QueueUrl' => QUEUE_URL,
            'MessageBody' => 'body ' . hash('sha256', time() . $i),
            'MessageGroupId' => 'php',
            // MessageBodyが重複せず、 「コンテンツに基づく重複排除」にチェックを入れた場合
            // MessageBodyの内容をsha256して自動で「MessageDeduplicationId」に設定してくれるので
            // MessageDeduplicationIdの指定は不要
            // MessageBodyが重複する場合は、MessageDeduplicationIdは必要
            // 'MessageBody' => 'body ' . $i,
            // 'MessageDeduplicationId' => hash('sha256', time() . $i),
        ]);
    }

} catch (\Aws\Sqs\Exception\SqsException $e) {
    var_export($e->getMessage());
}

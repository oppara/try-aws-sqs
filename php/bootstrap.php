<?php
require_once __DIR__ . '/vendor/autoload.php';

$dotenv = new Dotenv\Dotenv(__DIR__);
$dotenv->load();

define('QUEUE_URL', $_ENV['QUEUE_URL']);

function createClient()
{
    return new \Aws\Sqs\SqsClient([
        'credentials' => [
            'key' => $_ENV['AWS_KEY'],
            'secret' => $_ENV['AWS_SECRET_KEY'],
        ],
        'region' => 'us-west-2',
        'version'  => 'latest',
        // 'debug' => true,
    ]);

}





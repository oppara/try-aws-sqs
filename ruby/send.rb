require './bootstrap.rb'

def makeBody(idx)
  {
    'id' => SecureRandom.uuid,
    'body' => "body #{idx}"
  }.to_json
end

begin

  sqs = createClient

  5.times do |i|
    sqs.send_message(
      queue_url: ENV['QUEUE_URL'],
      message_body: makeBody(i),
      message_group_id: 'ruby'
    )
  end

rescue Aws::SQS::Errors::ServiceError => e
  p 'Error: ' << e.message
  exit(false)
end

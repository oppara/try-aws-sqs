require './bootstrap.rb'

begin

  sqs = createClient

  result = sqs.receive_message({
    queue_url: ENV['QUEUE_URL'],
    max_number_of_messages: 10
  })

  result.messages.each do |message|
    puts message.body

    sqs.delete_message({
      queue_url: ENV['QUEUE_URL'],
      receipt_handle: message.receipt_handle
    })
  end

rescue Aws::SQS::Errors::ServiceError => e
  p 'Error: ' << e.message
  exit(false)
end

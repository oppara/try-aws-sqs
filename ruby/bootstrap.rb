require 'aws-sdk'
require 'dotenv'

Dotenv.load

def createClient
  Aws::SQS::Client.new(
    access_key_id: ENV['AWS_KEY'],
    secret_access_key: ENV['AWS_SECRET_KEY'],
    region: 'us-west-2'
  )
end

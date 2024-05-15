wrk.method = "POST"
wrk.headers["Content-Type"] = "application/json"

-- Define multiple cases for different types of notifications with longer content messages
local cases = {
	'{"type":"email","content":"This is a test email notification with a longer content message to test the system."}',
	'{"type":"sms","content":"This is a test SMS notification with a longer content message to test the system."}',
	'{"type":"push","content":"This is a test push notification with a longer content message to test the system."}',
	'{"type":"email","content":"Another email notification with even more content to see how the system handles larger payloads."}',
	'{"type":"sms","content":"Another SMS notification with even more content to see how the system handles larger payloads."}',
	'{"type":"push","content":"Another push notification with even more content to see how the system handles larger payloads."}',
}

-- Function to select a random case for each request
request = function()
	local case = cases[math.random(1, #cases)]
	wrk.body = case
	return wrk.format(nil, "/notify")
end

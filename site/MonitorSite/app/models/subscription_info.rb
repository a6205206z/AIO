class SubscriptionInfo < ActiveRecord::Base
	validates_presence_of	:service_id,																			:message => "service id can not be empty!"
	validates_presence_of	:Subscriber_name,																		:message => "Subscriber name can not be empty!"
	validates_presence_of	:Subscriber_mail_address,																:message => "Subscriber mail address can not be empty!"

	validates_format_of 	:Subscriber_mail_address, 	:with => /\A([^@\s]+)@((?:[-a-z0-9]+\.)+[a-z]{2,})\Z/i， 	:message => "maile address format is not correct!"
end
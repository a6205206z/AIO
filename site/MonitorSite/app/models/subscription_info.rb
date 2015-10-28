class SubscriptionInfo < ActiveRecord::Base
	validates_presence_of	:service_id,																			:message => "can not be empty!"
	validates_presence_of	:subscriber_name,																		:message => "can not be empty!"
	validates_presence_of	:subscriber_mail_address,																:message => "can not be empty!"

	validates_format_of 	:subscriber_mail_address, 	:with => /\A([^@\s]+)@((?:[-a-z0-9]+\.)+[a-z]{2,})\Z/i, 	:message => "format is not correct!"

	def saveAndUpdateTime
		self.subscribe_date = Time.now
		self.save()
	end

end
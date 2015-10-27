class Service < ActiveRecord::Base
	validates_presence_of	:real_address,  																		:message => "real address can not be empty!"
	validates_presence_of	:esb_address,  																			:message => "esb address can not be empty!"
	validates_presence_of	:name,  																				:message => "name can not be empty!"
	validates_presence_of	:register_user_mail_address,  															:message => "user maill address can not be empty!"
	validates_presence_of	:environment,  																			:message => "environment can not be empty!"
	validates_presence_of	:major_version,  																		:message => "major version can not be empty!"
	validates_presence_of	:minor_version,  																		:message => "minor version can not be empty!"

	validates_format_of 	:register_user_mail_address, :with => /\A([^@\s]+)@((?:[-a-z0-9]+\.)+[a-z]{2,})\Z/i,	:message => "maile address format is not correct!"

	def saveAndUpdateTime
		self.last_modify_date = Time.now
		self.save()
	end

	def self.loadService()
		all()
	end
end
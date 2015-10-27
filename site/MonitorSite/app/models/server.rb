class Server < ActiveRecord::Base
	validates_uniqueness_of :server_name,	:case_sensitive => false, 	:message => "already exists!"
	validates_presence_of	:server_name,  								:message => "can not be empty!"
	validates_presence_of	:ip_address,  								:message => "can not be empty!"
	validates_presence_of	:port,  									:message => "can not be empty!"
end
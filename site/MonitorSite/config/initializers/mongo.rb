MongoMapper.connection = Mongo::Connection.new('192.168.200.22',27017)
MongoMapper.database = "vesb"

if defined?(PhusionPassenger)
	PhusionPassenger.on_event(:starting_worker_process) do |forked|
		MongoMapper.connection.connect if forked
	end
end
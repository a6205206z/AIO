class Systemlog
	include MongoMapper::Document
	set_collection_name :systemlog

	key :server, String
	key :description, String
	key :processtime, Time
	key :requests, Integer
	key :proxyrequests, Integer
	key :cpuused, Float
	key :memused, Float
end
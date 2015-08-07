class Tracking
		include MongoMapper::Document
	set_collection_name :tracking

	key :appname, String
	key :location, String
	key :url, String
	key :clientip, String
	key :realclientip, String
	key :statuscode, Integer
	key :usetime, Integer
	key :invoketime, Time
end
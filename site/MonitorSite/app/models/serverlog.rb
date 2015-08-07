class Serverlog
	include MongoMapper::Document
	set_collection_name :serverlog

	key :server, String
	key :servername, String
	key :processtime, Time
	key :count, Integer
	key :maxusetime, Integer
	key :avgusetime, Integer
	key :maxesblogicusetime, Integer

	key :clientiplist, Array
end
class ServerMonitorController < ApplicationController
	def index
		
	end

	def server

	end


	def tracking_log
		@tracking = Tracking.where(:appname => params[:app], :usetime => {"$gte" => (params[:mintime].to_i*1000)}).order("invoketime desc").limit(50)
	end
end
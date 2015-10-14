class ServerMonitorController < ApplicationController
	def index
		
	end

	def server

	end


	def tracking_log
		@tracking = Tracking.where(appname: params[:app]).order("invoketime desc").limit(50)
	end
end
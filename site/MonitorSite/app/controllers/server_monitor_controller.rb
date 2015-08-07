class ServerMonitorController < ApplicationController
	def index
		
	end

	def server
		@systemlog = Systemlog.where(server: params[:server]).order("processtime desc").limit(200)
	end


	def tracking_log
		@serverlog = Serverlog.where(server: params[:server], servername: params[:app]).order("processtime desc").limit(50)
		@tracking = Tracking.where(appname: params[:app]).order("invoketime desc").limit(50)
	end
end
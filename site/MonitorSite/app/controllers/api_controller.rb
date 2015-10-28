class ApiController < ActionController::Base
	def server
		@systemlog = Systemlog.where(server: params[:server]).order("processtime desc").limit(params[:count]) 
	end

	def tracking_log
		#@serverlog = Serverlog.where(server: params[:server], servername: params[:app]).order("processtime desc").limit(params[:count])
		@serverlog = Serverlog.where(servername: params[:app]).order("processtime desc").limit(params[:count])
	end
end
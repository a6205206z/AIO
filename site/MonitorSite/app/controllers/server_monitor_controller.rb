class ServerMonitorController < ApplicationController
	def index
		@systemlog = Systemlog.where(server: params[:server]).order("processtime desc").limit(200)
	end
end
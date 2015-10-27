class ServerController < ApplicationController
	protect_from_forgery :except => :create

	def index
		@servers = Server.all().order("server_name")
	end

	def new

	end

	def edit
		@server = Server.find(params[:id])
	end

	def update
		@server = Server.find(params[:id])
		@server.update_attributes(server_params)
	end

	def create
		@server = Server.new(server_params)
		@server.save()
	end

	private
	def server_params  
		params.require(:server).permit(:id, :server_name, :description, :ip_address, :port, :host_name)
    end  
end
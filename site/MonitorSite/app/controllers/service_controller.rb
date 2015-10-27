class ServiceController < ApplicationController
	def new

	end

	def create
		@service = Service.new(service_params)
		@service.saveAndUpdateTime()
	end

	def index
		@services = Service.loadService()
	end

	private
	def service_params  
		params.require(:service).permit(:id, :real_address, :esb_address, :name, :register_user_mail_address, :environment, :major_version, :minor_version, :description)
    end  
end
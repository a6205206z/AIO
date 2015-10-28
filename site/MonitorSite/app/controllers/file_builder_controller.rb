class FileBuilderController < ActionController::Base

	def deployments
		if params[:service_name] == "all"
			@services = Service.all()
		else
			@services = Service.where(:name => params[:service_name])
		end
		render content_type: "text/plain"
	end
end
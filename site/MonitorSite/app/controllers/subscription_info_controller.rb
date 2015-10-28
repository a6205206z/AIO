class SubscriptionInfoController < ApplicationController
	def index
		@subscription_info = SubscriptionInfo.all()
	end

	def new

	end

	def create
		@subscription_info = SubscriptionInfo.new(subscription_info_params)
		@subscription_info.saveAndUpdateTime()
	end

	def subscription_info_params  
		params.require(:subscription_info).permit(:service_id, :subscriber_name, :subscriber_group_name, :description, :subscriber_mail_address)
    end  
end
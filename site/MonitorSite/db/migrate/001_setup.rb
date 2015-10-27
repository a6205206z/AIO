class Setup < ActiveRecord::Migration
	def self.up
		create_table :servers do |t|
			t.column :server_name, :string, :limit => 30, :default => "", :null => false, :unique => true
			t.column :description, :string, :limit => 200
			t.column :ip_address, :string, :default => "", :null =>false
			t.column :port, :integer, :default => 80, :null => false
			t.column :host_name, :string, :default => "", :null =>false
		end

		create_table :services do |t|
			t.column :real_address, :string, :default => "", :null => false
			t.column :esb_address, :string, :default => "", :null => false
			t.column :name, :string, :default => "", :null => false
			t.column :description, :string, :limit => 200
			t.column :register_user_mail_address, :string, :default => "", :null => false
			t.column :last_modify_date, :datetime, :null => false
			t.column :environment, :string, :default => "production" ,:null => false
			t.column :major_version, :integer, :default => 0, :null => false
			t.column :minor_version, :integer, :default => 0, :null => false
		end

		create_table :subscription_infos do |t|
			t.column :service_id, :integer, :null => false
			t.column :Subscriber_name, :string, :default => "", :null => false
			t.column :Subscriber_group_name, :string, :default => "", :null => false
			t.column :description, :string, :limit => 200
			t.column :Subscriber_mail_address, :string, :default => "", :null => false
			t.column :Subscribe_date, :datetime, :null => false
		end
	end
end
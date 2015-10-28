# encoding: UTF-8
# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# Note that this schema.rb definition is the authoritative source for your
# database schema. If you need to create the application database on another
# system, you should be using db:schema:load, not running all the migrations
# from scratch. The latter is a flawed and unsustainable approach (the more migrations
# you'll amass, the slower it'll run and the greater likelihood for issues).
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 1) do

  create_table "servers", force: :cascade do |t|
    t.string  "server_name", limit: 30,  default: "", null: false
    t.string  "description", limit: 200
    t.string  "ip_address",  limit: 255, default: "", null: false
    t.integer "port",        limit: 4,   default: 80, null: false
    t.string  "host_name",   limit: 255, default: "", null: false
  end

  create_table "services", force: :cascade do |t|
    t.string   "real_address",               limit: 255, default: "",           null: false
    t.string   "esb_address",                limit: 255, default: "",           null: false
    t.string   "name",                       limit: 255, default: "",           null: false
    t.string   "description",                limit: 200
    t.string   "register_user_mail_address", limit: 255, default: "",           null: false
    t.datetime "last_modify_date",                                              null: false
    t.string   "environment",                limit: 255, default: "production", null: false
    t.integer  "major_version",              limit: 4,   default: 0,            null: false
    t.integer  "minor_version",              limit: 4,   default: 0,            null: false
  end

  create_table "subscription_infos", force: :cascade do |t|
    t.integer  "service_id",              limit: 4,                null: false
    t.string   "subscriber_name",         limit: 255, default: "", null: false
    t.string   "subscriber_group_name",   limit: 255, default: "", null: false
    t.string   "description",             limit: 200
    t.string   "subscriber_mail_address", limit: 255, default: "", null: false
    t.datetime "subscribe_date",                                   null: false
  end

end

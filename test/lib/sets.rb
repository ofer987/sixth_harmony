require 'pry'

sets_dir = File.expand_path(File.join(File.dirname(__FILE__), 'sets'))

Dir.entries(sets_dir).select do |entry|
  entry =~ /\.rb$/
end.each do |file|
  require File.join(sets_dir, file).to_s
end

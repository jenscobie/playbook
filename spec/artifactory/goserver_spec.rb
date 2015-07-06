require 'spec_helper'

describe command('java -version') do
  its(:exit_status) { should eq 0 }
  its(:stdout) { should match /1.8/ }
end

describe port(8153) do
  it { should be_listening }
end
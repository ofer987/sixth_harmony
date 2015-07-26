require_relative '../spec_helper.rb'

include Sets

RSpec.describe MightySet do
  sets_and_sums =
    [
      [ [0, 1, 2, 3, 4, 5], 15 ],
      [ [-1, 0, 1, 2, 3, 4, 5], 15 ],
      [ [-1, 0, 1, 2, 3, 4, -1], 10 ],
      [ [-1, -10, 0, 1, 2, 3, 4, -1, -20], 10 ],
      [ [-1, -10, 0, 1, 2, 3, 4, -1, -20, 100], 100 ],
      [ [-1, -10, 0, 1, 2, 3, 4, -1, -20, 100, 120], 220 ],
      [ [-1, -10, 0, 1, 1000, 3, 4, -1, -20, 100, 120], 1207 ],
      [ [-1, -10, 0, 1, 1000, 3, 4, -1, -20, 100, 120, -2000, 2], 1207 ],
      [ [3000, -1, -10, 0, 1, 1000, 3, 4, -1, -20, 100, 120, -2000, 2], 4196 ],
      [ [989, -990, 0, 1000, -2000, 995], 1000 ],
      [ [0, 0, 0, 0], 0 ],
      [ [10, -10, 20, 0], 20 ],
    ]

  context 'The greatest sum of the subset of' do
    sets_and_sums.each do |set_and_sum|
      numbers = set_and_sum[0]
      expected_sum = set_and_sum[1]
      it "#{numbers} is #{expected_sum}" do
        actual_sum = (MightySet.new numbers).largest_sum
        expect(actual_sum).to eq(expected_sum)
      end
    end
  end
end

# frozen_string_literal: true

# Day 11
module Day11
  def self.part1
    nums = File.new('day11/input.txt').readline(chomp: true).split(' ').map(&:to_i)

    25.times { nums = blink(nums) }

    puts(nums.length)
  end

  def self.blink(nums)
    nums.flat_map { |num| individual_blink(num) }
  end

  def self.individual_blink(num)
    if num == 0
      1
    elsif even_digits?(num)
      half_zeroes = 10.pow(num_digits(num) / 2)
      first_half = num / half_zeroes
      second_half = num - first_half * half_zeroes

      [first_half, second_half]
    else
      num * 2024
    end
  end

  def self.num_digits(num)
    Math.log10(num).floor + 1
  end

  def self.even_digits?(num)
    num_digits(num).even?
  end

  private_class_method(:blink)
  private_class_method(:individual_blink)
  private_class_method(:num_digits)
  private_class_method(:even_digits?)
end

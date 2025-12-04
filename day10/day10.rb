# frozen_string_literal: true

module Day10
  def self.part1
    matrix = File.foreach('day10/input.txt').map { |line| line.strip.chars.map(&:to_i) }

    trailhead_score_sum = 0

    matrix.each_with_index do |r, row|
      r.each_with_index do |_, col|
        trailhead_score_sum += trailhead_score(matrix, row, col)
      end
    end

    trailhead_score_sum
  end

  def self.part2
    matrix = File.foreach('day10/input.txt').map { |line| line.strip.chars.map(&:to_i) }

    trailhead_score_sum = 0

    matrix.each_with_index do |r, row|
      r.each_with_index do |_, col|
        trailhead_score_sum += trailhead_score_rating(matrix, row, col)
      end
    end

    trailhead_score_sum
  end

  def self.trailhead_score(matrix, row, col)
    return 0 if matrix[row][col] != 0

    dfs_score_once(deep_copy_matrix(matrix), row, col)
  end

  def self.dfs_score_once(matrix, row, col)
    height = matrix[row][col]

    if height == 9
      # mark as visited for the current trailhead
      matrix[row][col] = -1
      return 1
    end

    up = row > 0 && matrix[row - 1][col] == height + 1 ?                    dfs_score_once(matrix, row - 1, col) : 0
    down = row < matrix.length - 1 && matrix[row + 1][col] == height + 1 ?  dfs_score_once(matrix, row + 1, col) : 0
    right = col < matrix.length - 1 && matrix[row][col + 1] == height + 1 ? dfs_score_once(matrix, row, col + 1) : 0
    left = col > 0 && matrix[row][col - 1] == height + 1 ?                  dfs_score_once(matrix, row, col - 1) : 0

    down + right + up + left
  end

  def self.trailhead_score_rating(matrix, row, col)
    return 0 if matrix[row][col] != 0

    dfs_score_all(matrix, row, col)
  end

  def self.dfs_score_all(matrix, row, col)
    height = matrix[row][col]

    return 1 if height == 9

    up = row > 0 && matrix[row - 1][col] == height + 1 ?                    dfs_score_all(matrix, row - 1, col) : 0
    down = row < matrix.length - 1 && matrix[row + 1][col] == height + 1 ?  dfs_score_all(matrix, row + 1, col) : 0
    right = col < matrix.length - 1 && matrix[row][col + 1] == height + 1 ? dfs_score_all(matrix, row, col + 1) : 0
    left = col > 0 && matrix[row][col - 1] == height + 1 ?                  dfs_score_all(matrix, row, col - 1) : 0

    down + right + up + left
  end

  def self.deep_copy_matrix(matrix)
    Marshal.load(Marshal.dump(matrix))
  end

  private_class_method(:trailhead_score)
  private_class_method(:trailhead_score_rating)
  private_class_method(:dfs_score_once)
  private_class_method(:dfs_score_all)
  private_class_method(:deep_copy_matrix)
end

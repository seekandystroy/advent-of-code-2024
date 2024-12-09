# frozen_string_literal: true

# Day 12
module Day12
  def self.part1
    matrix = File.foreach('day12/input.txt').map { |line| line.strip.chars.prepend(-1).append(-1) }
    matrix.prepend(Array.new(matrix[0].length, -1))
    matrix.append(Array.new(matrix[0].length, -1))

    visited = matrix.map { |row| row.map { |_| false } }
    row = 1
    col = 1
    sum = 0

    while row < matrix.length - 1
      perimeter, area, new_row, new_col = traverse(matrix, row, col, 0, 0, visited)

      # puts("Perimeter #{perimeter}; Area #{area}; New: #{new_row}, #{new_col}")

      sum += perimeter * area
      row = new_row
      col = new_col
    end

    puts sum
  end

  def self.traverse(matrix, row, col, perimeter, area, visited)
    symbol = matrix[row][col]
    visited[row][col] = true # visited
    # puts("Visiting #{row}, #{col}")
    area += 1

    if matrix[row][col + 1] == symbol && !visited[row][col + 1]
      perimeter, area = traverse(matrix, row, col + 1, perimeter, area, visited)
    elsif matrix[row][col + 1] != symbol
      perimeter += 1
    end

    if matrix[row + 1][col] == symbol && !visited[row + 1][col]
      perimeter, area = traverse(matrix, row + 1, col, perimeter, area, visited)
    elsif matrix[row + 1][col] != symbol
      perimeter += 1
    end

    if matrix[row][col - 1] == symbol && !visited[row][col - 1]
      perimeter, area = traverse(matrix, row, col - 1, perimeter, area, visited)
    elsif matrix[row][col - 1] != symbol
      perimeter += 1
    end

    if matrix[row - 1][col] == symbol && !visited[row - 1][col]
      perimeter, area = traverse(matrix, row - 1, col, perimeter, area, visited)
    elsif matrix[row - 1][col] != symbol
      perimeter += 1
    end

    new_row, new_col = find_next(visited, row)

    [perimeter, area, new_row, new_col]
  end

  def self.find_next(visited, row)
    (row..visited.length - 2).each do |r|
      (1..visited.length - 2).each do |c|
        return [r, c] unless visited[r][c]
      end
    end

    [visited.length - 1, visited.length - 1]
  end

  private_class_method(:traverse)
  private_class_method(:find_next)
end

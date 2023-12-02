{:ok, contents} = File.read("data.txt")

numbers = %{
  one: 1,
  two: 2,
  three: 3,
  four: 4,
  five: 5,
  six: 6,
  seven: 7,
  eight: 8,
  nine: 9,
}

matcher = %{
  one: "",
  two: "",
  three: "",
  four: "",
  five: "",
  six: "",
  seven: "",
  eight: "",
  nine: "",
  result: -1
}


str = "1234"

# the accumulator is a tuple containing two values:
# 1 - an empty string containing the longest match so far
# 2 - The letter position to search against
res = Enum.reduce(String.graphemes(str), "", fn char, acc ->
  # First, if the accumulators string matches a key or is a number, return that number
  case char do
    Regex.match?(~r/\d/, char) -> "found a letter"
  end

end)

IO.puts(res)



# IO.puts(String.contains?(str, "nine"))



# lines = Enum.map(String.split(contents, "\r\n"), fn x ->

#   # # find the earliest instance of any number
#   earliestInstance = -1
#   latestInstance = -1

#   Enum.each(numbers, fn {key, value} ->
#     numberString = Atom.to_string(key)
#     String.split(x, numberString)
#     {:ok, val} = String.contains?(x, Atom)
#   end)

#   # String.starts_with?()

#   # lineDigits = String.replace(x, ~r/\D/, "")
#   # IO.puts(lineDigits)
#   # String.to_integer(String.first(lineDigits) <> String.last(lineDigits))
# end)

# IO.puts(lines)

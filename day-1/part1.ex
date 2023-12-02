{:ok, contents} = File.read("data.txt")

data = contents

lines = Enum.sum(String.split(data, "\r\n"), fn x ->
  lineDigits = String.replace(x, ~r/\D/, "")
  IO.puts(lineDigits)
  String.to_integer(String.first(lineDigits) <> String.last(lineDigits))
end )

IO.puts(lines)

module example.com/hello2

go 1.18

replace example.com/greetings2 => ../greetings

require example.com/greetings2 v0.0.0-00010101000000-000000000000

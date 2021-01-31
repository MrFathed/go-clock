package main

type placeholder [5]string

var zero placeholder = placeholder{
	"███",
	"█ █",
	"█ █",
	"█ █",
	"███",
}

var one placeholder = placeholder{
	"██ ",
	" █ ",
	" █ ",
	" █ ",
	"███",
}

var two placeholder = placeholder{
	"███",
	"  █",
	"███",
	"█  ",
	"███",
}

var three placeholder = placeholder{
	"███",
	"  █",
	"███",
	"  █",
	"███",
}

var four placeholder = placeholder{
	"█ █",
	"█ █",
	"███",
	"  █",
	"  █",
}

var five placeholder = placeholder{
	"███",
	"█  ",
	"███",
	"  █",
	"███",
}

var six placeholder = placeholder{
	"███",
	"█  ",
	"███",
	"█ █",
	"███",
}

var seven placeholder = placeholder{
	"███",
	"  █",
	"  █",
	"  █",
	"  █",
}

var eight placeholder = placeholder{
	"███",
	"█ █",
	"███",
	"█ █",
	"███",
}

var nine placeholder = placeholder{
	"███",
	"█ █",
	"███",
	"  █",
	"███",
}

var colon placeholder = placeholder{
	"   ",
	" █ ",
	"   ",
	" █ ",
	"   ",
}

var digits = [...]placeholder{
	zero, one, two, three, four, five, six, seven, eight, nine,
}

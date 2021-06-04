func interpret(command string) string {
	var res string
	l := len(command)
	for i, s := range command {
		if s == 'G' {
			res += string(s)
		} else if s == '(' && i+1 < l && command[i+1] == ')' {
			res += "o"
		} else if s == '(' && i+1 < l && command[i+1] == 'a' {
			res += "al"
		}
	}
	return res
}

func interpret(command string) string {
	command = strings.ReplaceAll(command, "(al)", "al")
	return strings.ReplaceAll(command, "()", "o")
}
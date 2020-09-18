package extract

// RegexList returns a list of predefined regex commands.
var RegexList = []PredefinedRegexCommand{
	{
		Name:    "email",
		Usage:   "returns emails",
		Regex:   `(?:[a-z0-9!#$%&'*+/=?^_\x60{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_\x60{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])`,
		Matches: []string{"hallo@hallo.de", "bl3456a@web.com", "ansdjabniudwgh@gmx.des"},
		Fails:   []string{"hallo@hallo", "hallo.de", "hallo@hallo@.de"},
	},
	{
		Name:    "ipv4",
		Usage:   "returns IP Version 4 addresses",
		Regex:   `\b(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\b`,
		Matches: []string{"255.255.255.255", "01.01.01.01", "1.1.1.1"},
		Fails:   []string{"256.1.1.1", "255.255.255.25v5"},
	},
	{
		Name:    "ipv6",
		Usage:   "returns IP Version 6 addresses",
		Regex:   `\b(?:[a-fA-F0-9]{1,4}:){7}[a-fA-F0-9]{1,4}\b`,
		Matches: []string{"1762:0:0:0:0:B03:1:AF18", "2001:0db8:85a3:0000:0000:8a2e:0370:7334", "2001:db8:85a3:0:0:8a2e:0:0"},
		Fails:   []string{"1762:0:0:0:0:B03:1:GF18", "2001:0db8:85a3:0000:0000:8a2e:0370:fffs", "2001:a6db8:85a3:0:0:8a2eg:0"},
	},
	{
		Name:    "ipaddress",
		Usage:   "returns IP addresses",
		Aliases: []string{"ip"},
		Regex:   `\b(((?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?))|(?:[a-fA-F0-9]{1,4}:){7}[a-fA-F0-9]{1,4})\b`,
		Matches: []string{"1762:0:0:0:0:B03:1:AF18", "2001:0db8:85a3:0000:0000:8a2e:0370:7334", "2001:db8:85a3:0:0:8a2e:0:0", "255.255.255.255", "01.01.01.01", "1.1.1.1"},
		Fails:   []string{"1762:0:0:0:0:B03:1:GF18", "2001:0db8:85a3:0000:0000:8a2e:0370:fffs", "2001:a6db8:85a3:0:0:8a2eg:0", "256.1.1.1", "255.255.255.25v5"},
	},
	{
		Name:    "url",
		Usage:   "returns url",
		Regex:   `(?i)(http://|https://|www.)([-a-zA-Z0-9]{1,63}\.[a-zA-Z0-9()]{1,6}\b)([-a-zA-Z0-9()@:%_\+.~#?&//=]*)?`,
		Matches: []string{"https://youtube.com", "https://sub.domain.com/asdasd/asdaasd/asdasd.html", "https://user-google.com"},
		Fails:   []string{"https://youtu#be.com", "https://s$ub.domain.com/asdasd/asdaasd/asdasd.html", "sdfasegsdg.aasd", "https://user@google.com", "https://user~google.com"},
	},
}

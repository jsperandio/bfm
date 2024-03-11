package model

type ListItem struct {
	Index       int
	Text        string
	Description string
	Short       rune
	Selected    func()
}

package src

var DefaultBoard = Board{
	NewTile(Red, Yellow, Green, Blue),
	NewTile(Blue, Pink, Green, Yellow),
	NewTile(Red, Yellow, Pink, Green),
	NewTile(Pink, Red, Blue, Yellow),
	NewTile(Green, Blue, Pink, Yellow),
	NewTile(Pink, Blue, Red, Yellow),
	NewTile(Blue, Pink, Red, Green),
	NewTile(Pink, Red, Green, Yellow),
	NewTile(Red, Blue, Pink, Green),
}

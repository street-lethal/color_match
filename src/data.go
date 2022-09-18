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

var DefaultVectoredBoard = Board{
	NewVectoredTile(Green, Yellow, Red, Blue, false, false, true, true),
	NewVectoredTile(Red, Yellow, Green, Blue, false, false, true, true),
	NewVectoredTile(Green, Blue, Yellow, Red, false, false, true, true),
	NewVectoredTile(Blue, Red, Yellow, Green, false, false, true, true),
	NewVectoredTile(Green, Blue, Red, Yellow, false, false, true, true),
	NewVectoredTile(Red, Blue, Green, Yellow, false, false, true, true),
	NewVectoredTile(Blue, Green, Red, Yellow, false, false, true, true),
	NewVectoredTile(Yellow, Red, Blue, Green, false, false, true, true),
	NewVectoredTile(Yellow, Green, Red, Blue, false, false, true, true),
}

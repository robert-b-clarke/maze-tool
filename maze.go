package main

import (
	"flag"
	"fmt"
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
	"math/rand"
	"strings"
	"time"
)

type cell struct {
	row   int
	col   int
	north *cell
	south *cell
	east  *cell
	west  *cell
	links []*cell
}

// Create a birectional link to another cell
func (c *cell) linkTo(neighbour *cell) {
	c.links = append(c.links, neighbour)
	neighbour.links = append(neighbour.links, c)
}

// Determine if a cell has a link to another
func (c *cell) hasLinkTo(neighbour *cell) bool {
	for _, link := range c.links {
		if link == neighbour {
			return true
		}
	}
	return false
}

// Get a pointer to a neighbouring cell
func (c *cell) randomNeighbour() *cell {
	var neighbours []*cell
	if c.north != nil {
		neighbours = append(neighbours, c.north)
	}
	if c.south != nil {
		neighbours = append(neighbours, c.south)
	}
	if c.east != nil {
		neighbours = append(neighbours, c.east)
	}
	if c.west != nil {
		neighbours = append(neighbours, c.west)
	}
	rand.Seed(time.Now().UnixNano())
	var neighbour *cell
	// start with a random cell
	neighbour = neighbours[rand.Intn(len(neighbours))]
	return neighbour
}

type grid struct {
	width  int
	height int
	cells  [][]cell
}

// Construct a new maze grid with no algorithm applied
func newGrid(width int, height int) grid {
	g := grid{}
	g.width = width
	g.height = height
	g.cells = make([][]cell, height)
	for r := 0; r < height; r++ {
		g.cells[r] = make([]cell, width)
	}
	// initialise all cells with sensible values
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			g.cells[r][c].row = r
			g.cells[r][c].col = c
			g.cells[r][c].links = nil
			g.cells[r][c].north = g.cellAt(c, r-1)
			g.cells[r][c].south = g.cellAt(c, r+1)
			g.cells[r][c].west = g.cellAt(c+1, r)
			g.cells[r][c].east = g.cellAt(c-1, r)
		}
	}
	return g
}

// Create a new grid object with the Aldous Broder maze algorithm
func newAldousBroderGrid(width int, height int) grid {
	// init a blank maze
	g := newGrid(width, height)
	rand.Seed(time.Now().UnixNano())
	// start with a random cell
	var neighbour *cell
	var currentCell *cell
	currentCell = &(g.cells[rand.Intn(g.height)][rand.Intn(g.width)])
	unvisited := (g.width * g.height) - 1
	for unvisited > 0 {
		neighbour = currentCell.randomNeighbour()
		if len(neighbour.links) == 0 {
			currentCell.linkTo(neighbour)
			unvisited--
		}
		currentCell = neighbour
	}
	return g
}

// Create a new grid object with the Binary Tree maze algorithm
func newBinaryTreeGrid(width int, height int) grid {
	// init a blank maze
	g := newGrid(width, height)
	rand.Seed(time.Now().UnixNano())
	for r := 0; r < g.height; r++ {
		for c := 0; c < g.width; c++ {
			if g.cells[r][c].north != nil && g.cells[r][c].east != nil {
				if rand.Intn(2) == 1 {
					g.cells[r][c].linkTo(g.cells[r][c].east)
				} else {
					g.cells[r][c].linkTo(g.cells[r][c].north)
				}
			} else if g.cells[r][c].north != nil {
				g.cells[r][c].linkTo(g.cells[r][c].north)
			} else if g.cells[r][c].east != nil {
				g.cells[r][c].linkTo(g.cells[r][c].east)
			}
		}
	}
	return g
}

// Get a pointer to a cell at a given grid location
func (g *grid) cellAt(x int, y int) *cell {
	if y < 0 || y >= g.height {
		return nil
	}
	if x < 0 || x >= g.width {
		return nil
	}
	return &g.cells[y][x]
}

// Render a maze grid as a PNG image
func (g *grid) gridToPng(filename string, cellSize int, border int) {
	// TODO - some messy dupe in here - consider refactor to dedicated canvas object
	canvasWidth := cellSize*g.width + border
	canvasHeight := cellSize*g.height + border
	// create canvas
	dest := image.NewRGBA(image.Rect(0, 0, canvasWidth, canvasHeight))
	gcon := draw2dimg.NewGraphicContext(dest)

	// Set pen properties
	gcon.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	gcon.SetLineWidth(float64(border))

	// draw cell boundaries
	for _, row := range g.cells {
		for _, cell := range row {
			topLeftX := float64(cell.col*cellSize) + 1.0
			topLeftY := float64(cell.row*cellSize) + 1.0
			if !cell.hasLinkTo(cell.north) {
				gcon.MoveTo(topLeftX, topLeftY)
				gcon.LineTo(topLeftX+float64(cellSize), topLeftY)
				gcon.Close()
				gcon.FillStroke()
			}
			if !cell.hasLinkTo(cell.east) {
				gcon.MoveTo(topLeftX, topLeftY)
				gcon.LineTo(topLeftX, topLeftY+float64(cellSize))
				gcon.Close()
				gcon.FillStroke()
			}
		}
	}
	// Draw Southern + Western Borders
	gcon.MoveTo(1.0, float64(canvasHeight))
	gcon.LineTo(float64(canvasWidth), float64(canvasHeight)-1.0)
	gcon.MoveTo(float64(canvasWidth)-1.0, float64(canvasHeight)-1.0)
	gcon.LineTo(float64(canvasWidth)-1.0, 1.0)
	gcon.Close()
	gcon.FillStroke()
	draw2dimg.SaveToPngFile(filename, dest)
}

// Render a maze grid with ASCII art
func (g *grid) gridToascii() string {
	var sb strings.Builder
	for _, row := range g.cells {
		for _, cell := range row {
			if cell.hasLinkTo(cell.north) {
				sb.WriteString("+  ")
			} else {
				sb.WriteString("+--")
			}
		}
		sb.WriteString("+\n")
		for _, cell := range row {
			if cell.hasLinkTo(cell.east) {
				sb.WriteString("   ")
			} else {
				sb.WriteString("|  ")
			}
		}
		sb.WriteString("|\n")
	}
	for c := 0; c < g.width; c++ {
		sb.WriteString("+--")
	}
	sb.WriteString("+\n")

	return sb.String()
}

func main() {
	// Parse cmdline args
	widthPtr := flag.Int("width", 5, "width of the maze, in cells")
	heightPtr := flag.Int("height", 5, "height of the maze, in cells")
	algoPtr := flag.String("algorithm", "", "aldousbroder | binarytree")
	pngFilePtr := flag.String("pngfile", "", "location of a file to output png")
	flag.Parse()
	// define grid
	var my_grid grid
	// apply algorithm
	switch *algoPtr {
	case "aldousbroder":
		my_grid = newAldousBroderGrid(*widthPtr, *heightPtr)
	case "binarytree":
		my_grid = newBinaryTreeGrid(*widthPtr, *heightPtr)
	default:
		fmt.Println("No maze algorithm - using default grid")
		my_grid = newGrid(*widthPtr, *heightPtr)
	}
	fmt.Println(my_grid.gridToascii())
	if len(*pngFilePtr) > 0 {
		my_grid.gridToPng(*pngFilePtr, 30, 2)
	}
}

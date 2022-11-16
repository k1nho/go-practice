package main


func pix(x, y, mode int) int {
  switch mode {
    case 1:
      return x^y
    case 2:
      return (x+y)/2
    default:
      return x*y
  }
}

func Pic(dx, dy int) [][]uint8 {
  grid := make([][]uint8 , dy)  

  for i:= range grid {
    grid[i] = make([]uint8, dx)
    for j:= range grid[i] {
      grid[i][j] = uint8(pix(i, j, 1))
    }
  }

  return grid;
}

func main() {
  pic.Show(Pic)
}

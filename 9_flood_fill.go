// Problem desc: https://leetcode.com/problems/flood-fill/

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    startingColor := image[sr][sc] 
    if startingColor == newColor {
        return image
    }

    var dfs func(row, col int) 
    dfs = func(row, col int) {
        if row < 0 || row >= len(image) || col < 0 || col >= len(image[0]) || startingColor != image[row][col] {
            return 
        }

        image[row][col] = newColor

        dfs(row-1, col)
        dfs(row+1, col)
        dfs(row, col-1)
        dfs(row, col+1)
    }

    dfs(sr, sc) 
    return image
}

package uitest

import (
	"fmt"
	"os"
	"path/filepath"
)

func generateByteFile(name string, size int) string {
	// convert size from MB to bytes
	size *= 1000000
	data := make([]byte, size)

	filePath := fmt.Sprintf("./testdata/%s", name)
	f, err := os.Create(filePath)

	if err != nil {
		panic(err)
	}

	_, err = f.Write(data)

	if err != nil {
		panic(err)
	}

	file, err := filepath.Abs(filePath)

	if err != nil {
		panic(err)
	}

	err = f.Close()

	if err != nil {
		panic(err)
	}

	return file
}

func removeFiles(files []string) {
	for _, file := range files {
		err := os.Remove(fmt.Sprintf("./testdata/%s", file))

		if err != nil {
			panic(err)
		}
	}
}

// func generateEmptyFile(t *testing.T, ctx *testcontext.Context, name string, size memory.Size) string {
// 	path := ctx.File(name)
// 	f, err := os.Create(path)
// 	require.NoError(t, err)
// 	defer func() { require.NoError(t, f.Close()) }()
// 	require.NoError(t, f.Truncate(size.Int64()))
// 	return path
// }

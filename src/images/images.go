package images

import ( 
    "image/png"
    "github.com/nfnt/resize"
    "Response"
    "fmt"
    "os"
)

func Resize(name string) {

	file, err := os.Open(name)
	if err != nil {
		Response.Print(fmt.Sprintf("%s\n", err))
	}

	img, err := png.Decode(file)
	if err != nil {
		Response.Print(fmt.Sprintf("%s\n", err))
	}
	file.Close()

	m := resize.Resize(500, 0, img, resize.Lanczos3)

	out, err := os.Create(name)
	if err != nil {
		Response.Print(fmt.Sprintf("%s\n", err))
	}
	defer out.Close()
	png.Encode(out, m)
}
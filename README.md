# Mock Data for crud api practice

## Description

Generates mock data for AI Generated Images

## Struct of the data

```bash
type GeneratedImage struct {
	ID         string  `json:"id"`
	Title      string  `json:"title"`
	Resolution string  `json:"resolution"`
	Model      string  `json:"model"`
	Prompt     string  `json:"prompt"`
	URL        string  `json:"url"`
	Author     *Author `json:"author"`
	Likes      int     `json:"likes"`
	Views      int     `json:"views"`
	CreatedAt  string  `json:"created_at"`
}

type Author struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	ImagesNum int    `json:"images_num"`
	CreatedAt string `json:"created_at"`
}
```

## How to run

```bash
go run *.go
```

## Author

- **Vadim Egorov** - [My GitHub](https://github.com/vadimegorov13)

# htmlcapture

## 📸 Overview

**htmlcapture** is a Go package that captures high-quality screenshots from URLs, HTML files, or raw HTML strings. It is optimized for Instagram post sizes and supports dynamic HTML rendering.

---

## 🚀 Installation

To install **htmlcapture**, use:

```sh
 go get github.com/rohankarn35/htmlcapture
```

---

## 📌 Features

- ✅ Capture from **URL, HTML file, or raw HTML string**
- ✅ Supports **dynamic HTML rendering** with external variables
- ✅ Capture **specific elements** using **CSS selectors**
- ✅ Default **Instagram-optimized** screenshot sizes
- ✅ Uses **headless Chrome (chromedp)** for rendering

---

## 🛠 Usage

### 1️⃣ **Basic Usage: Capture a Raw HTML String**

```go
import (
    "fmt"
    "os"
    "github.com/rohankarn35/htmlcapture"
)

func main() {
    opts := htmlcapture.CaptureOptions{
        Input: `<html><body><h1>Hello, World!</h1></body></html>`,
    }
    img, err := htmlcapture.Capture(opts)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    os.WriteFile("screenshot.png", img, 0644)
}
```

---

### 2️⃣ **Capture an HTML File (With Dynamic Variables)**

**Example: `template.html`**

```html
<html>
  <body>
    <h1>Welcome {{.User}}, to {{.Website}}!</h1>
  </body>
</html>
```

**Go Code:**

```go
opts := htmlcapture.CaptureOptions{
    Input: "template.html",
    Variables: map[string]string{
        "User": "Alice",
        "Website": "Wonderland",
    },
}
img, err := htmlcapture.Capture(opts)
```

---

### 3️⃣ **Capture a Website URL**

```go
opts := htmlcapture.CaptureOptions{
    Input: "https://example.com",
    ViewportW: 1920,
    ViewportH: 1080,
}
img, err := htmlcapture.Capture(opts)
```

---

### 4️⃣ **Capture a Specific Element (CSS Selector)**

```go
opts := htmlcapture.CaptureOptions{
    Input: "https://example.com",
    Selector: "#main-content",
}
img, err := htmlcapture.Capture(opts)
```

---

## 📝 Configuration Options

| Option      | Type              | Default  | Description                                   |
| ----------- | ----------------- | -------- | --------------------------------------------- |
| `Input`     | string            | Required | URL, HTML string, or file path                |
| `ViewportW` | int               | 1080     | Viewport width                                |
| `ViewportH` | int               | 1350     | Viewport height                               |
| `Selector`  | string            | ""       | CSS selector for capturing a specific element |
| `Variables` | map[string]string | `{}`     | Key-value pairs for dynamic HTML              |

---

## 📖 License

MIT License

---

## 💡 Future Enhancements

- ✅ Add **PDF output support** 📄
- ✅ Improve **performance optimizations** 🚀
- ✅ Additional **image quality controls** 🎨

**Contributions are welcome! Feel free to open issues or submit pull requests.** 😃

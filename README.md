## Control requirement justification structure
This file explains the structure of the individual control justification files

Yaml Format
```
---
control:
  a:
  - title: Title of control requirement justifications
    justifications:
    - text: Text justification
      image:
        text: Image text
        url: url
      link:
        text: URL text
        url: url
```

JSON format
```json
{
  "control": {
    "a": [
      {
        "title": "Title of control requirement justifications",
        "justifications": [
          {
            "text": "Text justification",
            "image": {"text": "Image text", "url": "url"},
            "link": {"text": "URL text", "url": "url"}
          }
        ]
      }
    ]
  }
}
```

Markdown Format
### Control
#### a
- ##### Title of control requirement justifications
  - Text justification
  - Image Text ![Image text](http://dummyimage.com/300x100/ffffff/131961.jpg&text=Image+Justification)
  - [URL text](https://18f.gsa.gov/)

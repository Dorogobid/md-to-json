# md-to-json

### Convert markdown file with quiz questions to JSON format

#### Input file structure

![md file](/assets//md-file.png)

#### Output file structure

```
[
  {
    "category_id": 0,
    "question_text": "What is this code an example of?",
    "code_block": "let val = (Double)6\n",
    "code_type": "swift",
    "answers": [
      { "answer_text": "a syntax issue", "is_true": true },
      { "answer_text": "typecasting", "is_true": false },
      { "answer_text": "assignment", "is_true": false },
      { "answer_text": "initialization", "is_true": false }
    ]
  }
]
```

#### Usage

```
md-to-json file-name.md
```

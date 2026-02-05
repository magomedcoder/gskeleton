## GUtil

Набор вспомогательных утилит для Go‑сервисов.

Репозиторий содержит небольшие, переиспользуемые обёртки и хелперы для типичных задач.

```bash
go get github.com/magomedcoder/gutil

import (
    "github.com/magomedcoder/gutil/ginutil"
    "github.com/magomedcoder/gutil/jwtutil"
)

token, err := jwtutil.GenerateToken("")
```

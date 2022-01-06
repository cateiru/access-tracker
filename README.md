[![codecov](https://codecov.io/gh/cateiru/access-tracker/branch/main/graph/badge.svg?token=XPDUL0ZJR2)](https://codecov.io/gh/cateiru/access-tracker)

# access-tracker

Access tracking

## How To Use

### Create tracking

```http
POST https://access-tracker-nwu7rw5hea-an.a.run.app/u
Content-Type: application/x-www-form-urlencoded

redirect=https://example.com
```

```text
{
  "track_id": "fdd6b450fd",
  "access_key": "88664dd2981526f89895ddb37c313f4b59f83dbc57591d3ff9e6eaeb9d3a408f",
  "redirect_url": "https://example.com"
}
```

### Access Tracking url

```http
GET https://access-tracker-nwu7rw5hea-an.a.run.app/[id]/
```

### Show Access Histories

```http
GET https://access-tracker-nwu7rw5hea-an.a.run.app/u?id=[id]&key=[key]
```

```text
[
  {
    "unique_id": "63dbd5e1d954937fa420910aa3f7ac3af6bff1f8031279fc825c76c6bbb733c3",
    "ip": "172.0.0.1",
    "time": "2021-10-27T05:06:51.916374Z"
  }
]
```

### Delete Tracking url

```http
DELETE https://access-tracker-nwu7rw5hea-an.a.run.app/u?id=[id]&key=[key]
```

## LICENSE

[MIT](./LICENSE)

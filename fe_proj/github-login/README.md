



## Installation

```bash
$ npm install
```

## Running the app

```bash
# development
$ npm run start

# watch mode
$ npm run start:dev

# production mode
$ npm run start:prod
```

## Test

```bash
# unit tests
$ npm run test

# e2e tests
$ npm run test:e2e

# test coverage
$ npm run test:cov
```

## 操作

* step 1: 访问 http://localhost:3000/login  这个会访问github的授权页面
* step 2：同意授权过后，回调 http://localhost:3000/callback
返回的数据如下，一般取`id`,`username`,`email`啥的
```json
{
    "id": "16758423",
    "nodeId": "MDQ6VXNlcjE2NzU4NDIz",
    "displayName": null,
    "username": "vencent2006",
    "profileUrl": "https://github.com/vencent2006",
    "photos": [
        {
            "value": "https://avatars.githubusercontent.com/u/16758423?v=4"
        }
    ],
    "provider": "github",
    "_raw": "{\"login\":\"vencent2006\",\"id\":16758423,\"node_id\":\"MDQ6VXNlcjE2NzU4NDIz\",\"avatar_url\":\"https://avatars.githubusercontent.com/u/16758423?v=4\",\"gravatar_id\":\"\",\"url\":\"https://api.github.com/users/vencent2006\",\"html_url\":\"https://github.com/vencent2006\",\"followers_url\":\"https://api.github.com/users/vencent2006/followers\",\"following_url\":\"https://api.github.com/users/vencent2006/following{/other_user}\",\"gists_url\":\"https://api.github.com/users/vencent2006/gists{/gist_id}\",\"starred_url\":\"https://api.github.com/users/vencent2006/starred{/owner}{/repo}\",\"subscriptions_url\":\"https://api.github.com/users/vencent2006/subscriptions\",\"organizations_url\":\"https://api.github.com/users/vencent2006/orgs\",\"repos_url\":\"https://api.github.com/users/vencent2006/repos\",\"events_url\":\"https://api.github.com/users/vencent2006/events{/privacy}\",\"received_events_url\":\"https://api.github.com/users/vencent2006/received_events\",\"type\":\"User\",\"site_admin\":false,\"name\":null,\"company\":null,\"blog\":\"\",\"location\":null,\"email\":null,\"hireable\":null,\"bio\":\"be the one\",\"twitter_username\":null,\"notification_email\":null,\"public_repos\":51,\"public_gists\":0,\"followers\":3,\"following\":25,\"created_at\":\"2016-01-18T12:01:56Z\",\"updated_at\":\"2024-09-10T03:53:01Z\"}",
    "_json": {
        "login": "vencent2006",
        "id": 16758423,
        "node_id": "MDQ6VXNlcjE2NzU4NDIz",
        "avatar_url": "https://avatars.githubusercontent.com/u/16758423?v=4",
        "gravatar_id": "",
        "url": "https://api.github.com/users/vencent2006",
        "html_url": "https://github.com/vencent2006",
        "followers_url": "https://api.github.com/users/vencent2006/followers",
        "following_url": "https://api.github.com/users/vencent2006/following{/other_user}",
        "gists_url": "https://api.github.com/users/vencent2006/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/vencent2006/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/vencent2006/subscriptions",
        "organizations_url": "https://api.github.com/users/vencent2006/orgs",
        "repos_url": "https://api.github.com/users/vencent2006/repos",
        "events_url": "https://api.github.com/users/vencent2006/events{/privacy}",
        "received_events_url": "https://api.github.com/users/vencent2006/received_events",
        "type": "User",
        "site_admin": false,
        "name": null,
        "company": null,
        "blog": "",
        "location": null,
        "email": null,
        "hireable": null,
        "bio": "be the one",
        "twitter_username": null,
        "notification_email": null,
        "public_repos": 51,
        "public_gists": 0,
        "followers": 3,
        "following": 25,
        "created_at": "2016-01-18T12:01:56Z",
        "updated_at": "2024-09-10T03:53:01Z"
    }
}
```

# travel-blog

travel blog

## Images

https://unsplash.com/search/photos/budapest

TODOs

- read the theme readme
- add favicon
- optimize images!!!!!
- add back tags
- deal with https://www.peekaboo.travel ssl error
- add pre-push hook 'hugo' build


Markdown reference
------------------

https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet

## Email 
```
cd email template && npx http-server

visit in browser at http://192.168.1.127:8080
```


to preview email:
```
cd email template &&  go run serve.go 
```

## REFRESH LONG LIVED 60 day token 
```
curl -i -X GET "https://graph.instagram.com/refresh_access_token
  ?grant_type=ig_refresh_token
  &access_token={long-lived-access-token}"
```

## GET POSTS 

```
set env var 
INSTAGRAM_TOKEN=
curl -X GET \    
´  'https://graph.instagram.com/me/media?fields=id,caption,media_type,media_url,permalink,thumbnail_url,timestamp,username&access_token='
```
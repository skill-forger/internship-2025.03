# Social Blogging - API Endpoints

### Profile

- `POST /api.myservice.com/v1/auth/sign-up` - Reader can sign up to become a blogger
- `POST /api.myservice.com/v1/auth/sign-in` - Blogger can sign in to the blogging website
- `GET /api.myservice.com/v1/profile/{:userId}` - Reader/Blogger can view other blogger profile information
- `PUT /api.myservice.com/v1/user/profile` - Blogger can update their profile information
- `PUT /api.myservice.com/v1/user/profile/change-password` - Blogger can change their account password
- `GET /api.myservice.com/v1/auth/verify` - Blogger can verify their email address upon signing up

---

### Favourite

- `POST /api.myservice.com/v1/user/favorites/bloggers/{:userId}` - Blogger can add another blogger into their following list
- `DELETE /api.myservice.com/v1/user/favorites/bloggers/{:userId}` - Blogger can remove blogger from their following list
- `GET /api.myservice.com/v1/user/favorites/bloggers` - Blogger can view all the bloggers from their following list
- `GET /api.myservice.com/v1/posts/bloggers/{:userId}` - Blogger can view all the posts of the following bloggers (em đang hiểu chức năng này là mọi người có thể xem tất cả bài viết của một blogger cụ thể)
- `POST /api.myservice.com/v1/user/favorites/posts/{:postId}` - Blogger can add a post into their favourite list
- `DELETE /api.myservice.com/v1/user/favorites/posts/{:postId}` - Blogger can remove a post from their favourite list
- `GET /api.myservice.com/v1/user/favorites/posts` - Blogger can view all posts from their favourite list

---

### Post

- `GET  /api.myservice.com/v1/posts` - Reader/Blogger can view all published blog posts, filter by specific condition
- `POST /api.myservice.com/v1/user/posts` - Blogger can create new blog post
- `PUT /api.myservice.com/v1/user/posts/{:postId}` - Blogger can edit current blog post
- `POST /api.myservice.com/v1/user/posts?isDraft = true/false` - Blogger can publish the blog post/ Blogger can draft the blog post to avoid incomplete blog post being published
- `PUT /api.myservice.com/v1/user/posts/{:postId}/publish` - Blogger can publish the blog post

---

### Comments

- `GET  /api.myservice.com/v1/posts/{:postId}/comments` - Reader/Blogger can view all comments in the blog posts
- `POST /api.myservice.com/v1/user/posts/{:postId}/comments` - Blogger can make a new comment/ Blogger can reply to another comment (with parentCMTID in request body)
- `PUT /api.myservice.com/v1/user/comments/{:commentId}` - Blogger can update their comment
- `DELETE /api.myservice.com/v1/user/comments/{:commentId}` - Blogger can delete their comment

---

### Tag

- `GET  /api.myservice.com/v1/tags` - Readers/Bloggers can view all blog tags
- `GET  /api.myservice.com/v1/posts/tags/{:tagId}` - Readers/Bloggers can view all blog posts belong to a particular tag
- `POST /api.myservice.com/v1/user/tags` - Blogger can create new blog tag
- `DELETE /api.myservice.com/v1/user/tags/{:tagId}` - Blogger can delete a tag that does not contain any blog

---

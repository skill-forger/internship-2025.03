# Social Blogging - API Endpoints

### Profile

- `POST /auth/sign-up` - Reader can sign up to become a blogger
- `POST /auth/sign-in` - Blogger can sign in to the blogging website
- `GET /profile/{:userId}` - Reader/Blogger can view other blogger profile information
- `PUT /profile` - Blogger can update their profile information
- `PUT /profile/change-password` - Blogger can change their account password
- `POST /auth/verify` - Blogger can verify their email address upon signing up

---

### Favourite

- `PUT /favorites/bloggers` - Blogger can add/remove blogger from their following list
- `GET /favorites/bloggers` - Blogger can view all the bloggers from their following list
- `GET /favorites/bloggers/posts` - Blogger can view all the posts of the following bloggers
- `PUT /favorites/posts` - Blogger can add/remove a post from their favourite list
- `GET /favorites/posts` - Blogger can view all posts from their favourite list

---

### Post

- `GET  /posts` - Reader/Blogger can view all published blog posts, filter by specific condition
- `GET /posts/{:postId}` Get detail of particular post
- `POST /posts` - Blogger can create new post (Default draft post)
- `PUT /posts/{:postId}` - Blogger can edit current blog post (publish<->draft)
- `DELETE /posts/{:postId}` - Blogger can delete the blog post

---

### Comments

- `GET  /comments?post_id` - Reader/Blogger can view all comments in the blog posts
- `POST /comments` - Blogger can make a new comment/ Blogger can reply to another comment (with parentCMTID in request body)
- `PUT /comments/{:commentId}` - Blogger can update their comment
- `DELETE /comments/{:commentId}` - Blogger can delete their comment

---

### Tag

- `GET  /tags` - Readers/Bloggers can view all blog tags
- `GET  /tags/{:tagId}/posts` - Readers/Bloggers can view all blog posts belong to a particular tag
- `POST /tags` - Blogger can create new blog tag
- `DELETE /tags/{:tagId}` - Blogger can delete a tag that does not contain any blog

---

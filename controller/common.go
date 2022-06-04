package controller

//公共返回对象包

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type VideoVo struct {
	Id            int64  `json:"id"`
	Author        UserVo `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

type UserVo struct {
	Id              int64  `json:"id"`
	Name            string `json:"name"`
	Avatar          string `json:"avatar"`
	Signature       string `json:"signature"`
	BackgroundImage string `json:"background_image"`
	FollowCount     int64  `json:"follow_count,omitempty"`
	FollowerCount   int64  `json:"follower_count,omitempty"`
	TotalFavorited  int64  `json:"total_favorited,omitempty"`
	FavoriteCount   int64  `json:"favorite_count,omitempty"`
	IsFollow        bool   `json:"is_follow"`
}

type CommentVo struct {
	Id         int64  `json:"id,omitempty"`
	User       UserVo `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

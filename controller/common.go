package controller

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

//type Video struct {
//	VideoId       int64  `json:"video_id"`
//	Author        User   `json:"author"`
//	PlayUrl       string `json:"play_url"`
//	CoverUrl      string `json:"cover_url"`
//	FavoriteCount int64  `json:"favorite_count"`
//	CommentCount  int64  `json:"comment_count"`
//	IsFavorite    bool   `json:"is_favorite"`
//	Title         string `json:"title"`
//}
//
//type User struct {
//	UserId        int64  `json:"user_id"`
//	Name          string `json:"name"`
//	FollowCount   int64  `json:"follow_count,omitempty"`
//	FollowerCount int64  `json:"follower_count,omitempty"`
//	IsFollow      bool   `json:"is_follow"`
//}

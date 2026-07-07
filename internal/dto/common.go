// 通用数据模型定义

package dto

// Response 统一 API 响应格式
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Paginated 分页响应数据
type Paginated struct {
	Items      interface{} `json:"items"`
	Total      int         `json:"total"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int         `json:"total_pages"`
}

// ArticleVO 资讯视图对象（给前端展示用）
type ArticleVO struct {
	ID           string     `json:"id"`
	Title        string     `json:"title"`
	Summary      string     `json:"summary"`
	Content      string     `json:"content,omitempty"`
	Category     CategoryVO `json:"category"`
	Author       string     `json:"author"`
	PublishDate  int64      `json:"publishDate"`   // UTC Unix 时间戳
	ReadTime     int        `json:"readTime"`
	XPReward     int        `json:"xpReward"`
	ImageURL     string     `json:"imageUrl"`
	Tags         []string   `json:"tags"`
	IsFeatured   bool       `json:"isFeatured"`
	ViewCount    int        `json:"viewCount"`
	CommentCount int        `json:"commentCount"`
	LikeCount    int        `json:"likeCount"`
	IsLiked      bool       `json:"isLiked"`       // 当前用户是否已点赞
	TimeAgo      string     `json:"timeAgo,omitempty"`    // 前端计算
	SourceName   string     `json:"sourceName,omitempty"`
	CollectedAt  int64      `json:"collectedAt,omitempty"` // UTC Unix 时间戳
}

// CategoryVO 分类视图对象
type CategoryVO struct {
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Color       string `json:"color"`
	Icon        string `json:"icon,omitempty"`
	ArticleCount int    `json:"articleCount,omitempty"`
}

// UserVO 用户视图对象
type UserVO struct {
	ID               string   `json:"id"`
	Email            string   `json:"email"`
	PlayerName       string   `json:"playerName"`
	Level            int      `json:"level"`
	XP               int      `json:"xp"`
	MaxXP            int      `json:"maxXp"`
	Title            string   `json:"title"`
	Tags             []string `json:"tags"`
	ReadArticles     int      `json:"readArticles"`
	LoginDays        int      `json:"loginDays"`
	AvatarURL        string   `json:"avatarUrl,omitempty"`
	ThemePreference  string   `json:"themePreference"`
	AchievementCount int      `json:"achievementCount"`
}

// CommentVO 评论视图对象
type CommentVO struct {
	ID        string `json:"id"`
	ArticleID string `json:"articleId"`
	Author    string `json:"author"`     // 用户昵称
	Avatar    string `json:"avatar"`     // 头像首字母
	Content   string `json:"content"`
	Time      int64  `json:"time"`       // UTC Unix 时间戳
	Likes     int    `json:"likes"`
	ParentID  string `json:"parentId,omitempty"`
}

// QuestVO 任务视图对象
type QuestVO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	XPReward    int    `json:"xpReward"`
	Progress    int    `json:"progress"`
	Target      int    `json:"target"`
	Status      string `json:"status"` // not_started / in_progress / completed
	Icon        string `json:"icon"`
}

// AchievementVO 成就视图对象
type AchievementVO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	UnlockedAt  int64  `json:"unlockedAt,omitempty"` // UTC Unix 时间戳
}

// TrendingTopicVO 热门话题视图对象
type TrendingTopicVO struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	ColorType string `json:"colorType"`
}

// SourceVO 数据源视图对象
type SourceVO struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	FeedURL       string `json:"feedUrl"`
	SiteURL       string `json:"siteUrl,omitempty"`
	CategoryID    string `json:"categoryId"`
	IsActive      bool   `json:"isActive"`
	FetchInterval int    `json:"fetchInterval"`
	LastFetchedAt int64  `json:"lastFetchedAt,omitempty"`
}

// ScheduleVO 调度配置视图对象
type ScheduleVO struct {
	ID             string `json:"id"`
	CategoryID     string `json:"categoryId"`
	CronExpression string `json:"cronExpression"`
	IsActive       bool   `json:"isActive"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=6"`
	RememberMe bool   `json:"rememberMe"`
}

// RegisterRequest 注册请求
type RegisterRequest struct {
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,eqfield=Password"`
	PlayerName      string `json:"playerName" binding:"required,min=2,max=20"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int    `json:"expiresIn"`
	User         UserVO `json:"user"`
}

// EmailVerifyRequest 邮箱验证请求
type EmailVerifyRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// EmailCodeVerifyRequest 验证码校验请求
type EmailCodeVerifyRequest struct {
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required,len=6"`
}

// ArticleListQuery 文章列表查询参数
type ArticleListQuery struct {
	Page      int    `form:"page" binding:"min=1"`
	PageSize  int    `form:"page_size" binding:"min=1,max=50"`
	Category  string `form:"category"`
	Tag       string `form:"tag"`
	Sort      string `form:"sort"`
	TimeRange string `form:"time_range"`
	Search    string `form:"search"`
	Featured  bool   `form:"featured"`
}

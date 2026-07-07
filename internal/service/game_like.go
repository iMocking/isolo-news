// 鐐硅禐绯荤粺涓氬姟閫昏緫锛堜娇鐢?Like 琛ㄧ鐞嗙湡瀹炴暟鎹級

package service

import (
	"context"
	"fmt"

	"isolo-news/internal/ent/like"
)

// ToggleLike 鍒囨崲鐐硅禐鐘舵€侊紝浣跨敤 Like 琛ㄧ鐞嗙湡瀹炴暟鎹?// 杩斿洖: isLiked锛堟搷浣滃悗鐨勭偣璧炵姸鎬侊級, error
func (s *GameService) ToggleLike(ctx context.Context, articleID, userID string) (bool, error) {
	// 妫€鏌ユ槸鍚﹀凡鐐硅禐
	exists, err := s.db.Like.Query().
		Where(like.UserIDEQ(userID), like.ArticleIDEQ(articleID)).
		Exist(ctx)
	if err != nil {
		return false, fmt.Errorf("鏌ヨ鐐硅禐鐘舵€佸け璐? %w", err)
	}

	if exists {
		// 鍙栨秷鐐硅禐锛氬垹闄?Like 璁板綍骞跺噺灏戞枃绔犵偣璧炴暟
		_, err = s.db.Like.Delete().
			Where(like.UserIDEQ(userID), like.ArticleIDEQ(articleID)).
			Exec(ctx)
		if err != nil {
			return false, fmt.Errorf("鍙栨秷鐐硅禐澶辫触: %w", err)
		}
		_, err = s.db.Article.UpdateOneID(articleID).AddLikeCount(-1).Save(ctx)
		if err != nil {
			return false, fmt.Errorf("鏇存柊鏂囩珷鐐硅禐鏁板け璐? %w", err)
		}
		return false, nil
	}

	// 娣诲姞鐐硅禐锛氬垱寤?Like 璁板綍骞跺鍔犳枃绔犵偣璧炴暟
	_, err = s.db.Like.Create().
		SetUserID(userID).
		SetArticleID(articleID).
		Save(ctx)
	if err != nil {
		return false, fmt.Errorf("鍒涘缓鐐硅禐澶辫触: %w", err)
	}
	_, err = s.db.Article.UpdateOneID(articleID).AddLikeCount(1).Save(ctx)
	if err != nil {
		return false, fmt.Errorf("鏇存柊鏂囩珷鐐硅禐鏁板け璐? %w", err)
	}
	return true, nil
}

// GetUserLikedMap 鎵归噺鏌ヨ鐢ㄦ埛瀵规寚瀹氭枃绔犵殑鐐硅禐鐘舵€?// 杩斿洖: map[articleID]isLiked
func (s *GameService) GetUserLikedMap(ctx context.Context, userID string, articleIDs []string) (map[string]bool, error) {
	if userID == "" || len(articleIDs) == 0 {
		return make(map[string]bool), nil
	}

	likes, err := s.db.Like.Query().
		Where(
			like.UserIDEQ(userID),
			like.ArticleIDIn(articleIDs...),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("鏌ヨ鐐硅禐鍒楄〃澶辫触: %w", err)
	}

	result := make(map[string]bool, len(articleIDs))
	for _, id := range articleIDs {
		result[id] = false
	}
	for _, l := range likes {
		result[l.ArticleID] = true
	}
	return result, nil
}

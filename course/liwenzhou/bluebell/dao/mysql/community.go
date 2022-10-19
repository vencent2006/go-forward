package mysql

import (
	"bluebell/models"
	"database/sql"

	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	if err = db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community is db")
			err = nil
		}
	}
	return
}

// GetCommunityDetailByID 跟进community_id查询社区详情
func GetCommunityDetailByID(cid int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := `select 
       	community_id, community_name, introduction, create_time, update_time 
		from community 
		where community_id = ?`
	if err = db.Get(community, sqlStr, cid); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("community not exist", zap.Int64("community_id", cid))
			err = ErrorInvalidID
		}
	}
	return
}

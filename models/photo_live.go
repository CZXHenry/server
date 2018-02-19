package models

import "github.com/sunyatsuntobee/server/logger"

// PhotoLive Model
type PhotoLive struct {
	ID                    int    `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	ExpectMembers         int    `xorm:"expect_members INT NOTNULL" json:"expect_members"`
	AdProgress            string `xorm:"ad_progress VARCHAR(100) NOTNULL" json:"ad_progress"`
	ActivityStageID       int    `xorm:"activity_stage_id INT INDEX(activity_stage_id_idx)" json:"activity_stage_id"`
	ManagerID             int    `xorm:"manager_id INT INDEX(manager_id_idx)" json:"manager_id"`
	PhotographerManagerID int    `xorm:"photographer_manager_id INT INDEX(photographer_manager_id_idx)" json:"photographer_manager_id"`
}

type PhotoLiveDataAccessObject struct{}

var PhotoLiveDAO *PhotoLiveDataAccessObject

func NewPhotoLive(expect_members int, ad_progress string, activity_stage_id int,
				  manager_id int, photographer_manager_id int) {
	return &PhotoLive{ExpectMembers:expect_members, AdProgress:ad_progress,
					  ActivityStageID:activity_stage_id, ManagerID:manager_id,
					  PhotographerManagerID:photographer_manager_id}
}

func (*PhotoLiveDataAccessObject) TableName() string {
	return "photo_lives"
}

func (*PhotoLiveDataAccessObject) InsertOne(photoLive *PhotoLive) {
	_, err := orm.Table(PhotoLiveDAO.TableName()).InsertOne(photoLive)
	logger.LogIfError(err)
}

func (*PhotoLiveDataAccessObject) FindAll() []PhotoLive {
	l := make([]PhotoLive, 0)
	err := orm.Table(PhotoLiveDAO.TableName()).Find(&l)
	logger.LogIfError(err)
	return l
}

func (*PhotoLiveDataAccessObject) UpdateOne(photoLive *PhotoLive) {
	_, err := orm.Table(PhotoLiveDAO.TableName()).ID(photoLive.ID).
		Update(photoLive)
	logger.LogIfError(err)
}

func (*PhotoLiveDataAccessObject) DeleteByID(id int) {
	var photoLive PhotoLive
	_, err := orm.Table(PhotoLiveDAO.TableName()).ID(id).Unscoped().
		Delete(&photoLive)
	logger.LogIfError(err)
}

func (*PhotoLiveDataAccessObject) FindFullAll() []PhotoLiveFull {
	l := make([]PhotoLiveFull, 0)
	err := orm.Table(PhotoLiveDAO.TableName()).
		Join("INNER", ActivityStageDAO.TableName(),
			"photo_lives.activity_stage_id=activity_stages.id").
		Join("INNER", ActivityDAO.TableName(),
			"activity_stages.activity_id=activities.id").
		Join("INNER", OrganizationDAO.TableName(),
			"activities.organization_id=organizations.id").
		Join("INNER", []string{UserDAO.TableName(), "u1"},
			"photo_lives.manager_id=u1.id").
		Join("INNER", []string{UserDAO.TableName(), "u2"},
			"photo_lives.photographer_manager_id=u2.id").
		Join("INNER", "photo_live_supervisor_relationships",
			"photo_lives.id=photo_live_supervisor_relationships.photo_live_id").
		Join("INNER", []string{UserDAO.TableName(), "u3"},
			"photo_live_supervisor_relationships.supervisor_id=u3.id").
		Find(&l)
	logger.LogIfError(err)
	return l
}

func (*PhotoLiveDataAccessObject) FindFullByID(id int) []PhotoLiveFull {
	l := make([]PhotoLiveFull, 0)
	err := orm.Table(PhotoLiveDAO.TableName()).
		Where("photo_lives.id=?", id).
		Join("INNER", ActivityStageDAO.TableName(),
			"photo_lives.activity_stage_id=activity_stages.id").
		Join("INNER", ActivityDAO.TableName(),
			"activity_stages.activity_id=activities.id").
		Join("INNER", OrganizationDAO.TableName(),
			"activities.organization_id=organizations.id").
		Join("INNER", []string{UserDAO.TableName(), "u1"},
			"photo_lives.manager_id=u1.id").
		Join("INNER", []string{UserDAO.TableName(), "u2"},
			"photo_lives.photographer_manager_id=u2.id").
		Join("INNER", "photo_live_supervisor_relationships",
			"photo_lives.id=photo_live_supervisor_relationships.photo_live_id").
		Join("INNER", []string{UserDAO.TableName(), "u3"},
			"photo_live_supervisor_relationships.supervisor_id=u3.id").
		Find(&l)
	logger.LogIfError(err)
	return l
}

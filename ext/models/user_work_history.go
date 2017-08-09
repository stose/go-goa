package models

import (
	"fmt"
	"github.com/hiromaily/go-goa/goa/app"
	"github.com/hiromaily/golibs/db/gorm"
)

// User is user object in Database
type UserWorkHistory struct {
	//Ctx *c.Ctx
	Db *gorm.GR
}

const TableName = "t_user_work_history"

//type Userworkhistory struct {
//	// Job Title
//	Title string `form:"title" json:"title" xml:"title"`
//	// Company name
//	Company string `form:"company" json:"company" xml:"company"`
//	// Country code
//	Country string `form:"country" json:"country" xml:"country"`
//	// worked period
//	Term *string `form:"term,omitempty" json:"term,omitempty" xml:"term,omitempty"`
//	// job description
//	Description *interface{} `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
//	// used techs
//	Techs *interface{} `form:"techs,omitempty" json:"techs,omitempty" xml:"techs,omitempty"`
//}
func (m *UserWorkHistory) GetUserWorks(userID int, userTechs *[]*app.Userworkhistory) error {
	sql := `
SELECT uwh.title, c.name as company, LOWER(mc.country_code) as country,
 CONCAT(DATE_FORMAT(IFNULL(uwh.started_at, ""),'%Y %b'), " - ", DATE_FORMAT(IFNULL(uwh.ended_at, ""),'%Y %b')) as term,
 uwh.description
 FROM t_user_work_history AS uwh
 LEFT JOIN t_company_detail AS cd ON uwh.company_branch_id = cd.id
 LEFT JOIN t_companies AS c ON c.id = cd.company_id
 LEFT JOIN m_countries AS mc ON mc.id = cd.country_id
 WHERE uwh.delete_flg=?
 AND cd.delete_flg=?
 AND c.delete_flg=?
 AND mc.delete_flg=?
 AND uwh.user_id=?
 ORDER BY uwh.started_at DESC
`
	//sql includes format character, so it can't be used below.
	//sql = fmt.Sprintf(sql, TableName)

	if err := m.Db.DB.Raw(sql, "0", "0", "0", "0", userID).Scan(userTechs).Error; err != nil {
		fmt.Println("[error]", err)
		return err
	}

	return nil
}

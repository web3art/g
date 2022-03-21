package apis

import (
	"encoding/json"
	"net/http"

	"github.com/web3art/g/internal/pkg/dao/model"
	"github.com/web3art/g/internal/pkg/db"
	"github.com/web3art/g/internal/pkg/thegraph/sword"
	"gorm.io/gorm"
)

func ReportUserStatics(w http.ResponseWriter, r *http.Request) {
	var userStatics model.UserStatistics

	if err := db.DB().Model(&model.UserStatistics{}).Where("id = 1").First(&userStatics).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if err := db.DB().Create(&model.UserStatistics{
				Id:            1,
				UserJoinCount: 1,
			}).Error; err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		userStatics.UserJoinCount++
		if err := db.DB().Save(&userStatics).Error; err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}
}

// export interface Statics {
//     plotLeft: number;
//     revisionsoFar: number;
//     userJoined: number;
// }
type Statics struct {
	PlotLeft      string `json:"plotLeft"`
	RevisionsoFar string `json:"revisionsoFar"`
	UserJoined    uint   `json:"userJoined"`
}

func GetStatics(w http.ResponseWriter, r *http.Request) {
	s := Statics{
		UserJoined:    1,
		RevisionsoFar: "40",
		PlotLeft:      "285",
	}

	var userStatics model.UserStatistics
	if err := db.DB().Model(&model.UserStatistics{}).Where("id = 1").First(&userStatics).Error; err == nil {
		s.UserJoined = userStatics.UserJoinCount
		// revisionsoFar = userStatics.RevisionsoFar
		// plotLeft = userStatics.PlotLeft
	}

	if q, err := sword.GetStatic(); err != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		// w.Write([]byte(err.Error()))
		// return
	} else {
		s.RevisionsoFar = q.Static.RevisionsoFar
		s.PlotLeft = q.Static.PlotLeft
	}

	if s.RevisionsoFar == "" {
		s.RevisionsoFar = "0"
	}

	if s.PlotLeft == "" {
		s.PlotLeft = "350"
	}

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(j)
}

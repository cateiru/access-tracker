package control

import (
	"github.com/yuto51942/access-tracker/types"
	"github.com/yuto51942/access-tracker/utils"
)

func Create() ([]byte, error) {
	id, err := utils.CreateId()
	if err != nil {
		return nil, err
	}

	types := types.New(types.Created{TrackId: id})

	return types.GetJson()
}

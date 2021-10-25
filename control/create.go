package control

import (
	"github.com/yuto51942/access-tracker/types"
	"github.com/yuto51942/access-tracker/utils"
)

func Create() (*types.Types, error) {
	id, err := utils.CreateId()
	if err != nil {
		return nil, err
	}

	return types.New(types.Created{TrackId: id}), nil
}
